package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// NOTE: Ordering of the functions represents the run order

var clients map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)
var broadcast chan Message = make(chan Message)

func registerWStoEndpoint(router *gin.Engine){
	router.GET("/ws", func(c *gin.Context){
		var upgrader websocket.Upgrader = getUpgrader()
		conn, err := upgrader.Upgrade(c.Writer,c.Request,nil)
		if err != nil {
			log.Println("Failed upgrading connection to WebSocket:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		log.Print("Successfully upgraded connection to WebSocket")
		defer conn.Close()

		clients[conn] = true
		log.Println("Registered client connection:")

		go listenForMessages(conn)
		broadcastToClients()
	})
}

func listenForMessages(conn *websocket.Conn){
	var db Database
	db.OpenConnection()
	defer db.CloseConnection()

	for {
		var temp map[string]string // keys: sender_id, username, contentRaw
		
		err := conn.ReadJSON(&temp) //blocking
		if err != nil {
			log.Println("Failed reading JSON from Socket:", err)
			delete(clients, conn)
			log.Println("Removing client connection:")
			return
		}

		var msg Message = createMessageHelper(temp, db)

		broadcast <- msg //blocking
	}
}

func createMessageHelper(temp map[string]string, db Database) Message {
	var msg Message
	var err error
	msg.Sender_id, err = strconv.Atoi(temp["sender_id"])
	if err != nil {
		log.Fatal("Failed converting string to int")
	}
	msg.ContentRaw = temp["contentRaw"]
	msg.ContentText = toContentText(temp["contentRaw"])
	msg.ContentMorse = toContentMorse(temp["contentRaw"])

	msg.Message_id, msg.Timestamp = db.PostMessage(
		msg.Sender_id,
		msg.ContentRaw,
		msg.ContentText,
		msg.ContentMorse,
	)

	return msg
}

func broadcastToClients() {
	for {
		msg := <-broadcast //blocking
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func getUpgrader() websocket.Upgrader {
	return websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true //Any origin is allowed
		},
	}
}

// **********************************************************************

func handleConnectionStdLib(
	w http.ResponseWriter,// <I> w/ res headers, body, status code 
	r *http.Request,      // Struct w/ req method, url, headers, query params, body
) {
	upgrader := getUpgrader()
	conn, err := upgrader.Upgrade(w,r,nil)
	if err != nil {
		log.Fatal("Failed upgrading connection to WebSocket:", err)
		return
	}
	log.Print("Successfully upgraded connection to WebSocket")
	defer conn.Close()

	clients[conn] = true
	log.Println("Registered client connection")

	go listenForMessages(conn)
	broadcastToClients()
}