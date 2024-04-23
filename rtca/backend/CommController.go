package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
	"strconv"
)

type CommController struct {
	clients map[*websocket.Conn]bool
	broadcast chan Message
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //Any origin is allowed
	},
}

func (cc *CommController) Start(){
	cc.init()

	go cc.broadcastToClients() //loops until error
	
	log.Println("Starting HTTP Server for ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil)) //loops until error
	
}

func (cc *CommController) init(){
	cc.clients = make(map[*websocket.Conn]bool) // key points to ws connection
	cc.broadcast = make(chan Message)
	log.Println("Allocated space for member variables")

	//Registers http request handler to a url path
	http.HandleFunc("/ws", cc.onClientConnect)
	log.Println("Registered request handler")
}

func (cc *CommController) onClientConnect(
	w http.ResponseWriter,// <I> w/ setters for http res headers, body, status code 
	r *http.Request,      // Struct w/ http req method, url, headers, query params, body
) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Failed to upgrade connection to WebSocket:", err)
		return
	} else {
		log.Print("Successfully upgraded connection to WebSocket")
	}
	defer conn.Close()

	cc.clients[conn] = true
	log.Println("Registered client's WebSocket Connection:")

	// TODO
			// Send client query of all Messages table

	cc.listenForClient(conn) // Forever looping
}

//Reading from client's Web Socket
func (cc *CommController) listenForClient(conn *websocket.Conn){
	var db AWS_RDS
	db.openConnection()
	defer db.closeConnection()

	for {
		var temp map[string]string // keys: sender_id, username, contentRaw
		
		err := conn.ReadJSON(&temp) //blocking
		if err != nil {
			log.Println("Failed reading JSON from Web Socket:", err)
			delete(cc.clients, conn)
			log.Println("UnRegistering client's WebSocket Connection:")
			return
		}

		var msg Message
		msg.Message_id = 0
		msg.Sender_id, err = strconv.Atoi(temp["sender_id"])
		if err != nil {
			log.Fatal("Failed to integer convert sender_id")
		}
		msg.ContentRaw = temp["contentRaw"]
		msg.ContentText = toContentText(temp["contentRaw"])
		msg.ContentMorse = toContentMorse(temp["contentRaw"])
		msg.Message_id, msg.Timestamp = db.addMessage(
			msg.Sender_id,
			msg.ContentRaw,
			msg.ContentText,
			msg.ContentMorse,
		)

		cc.broadcast <- msg //blocking
		log.Println("HTTP Server received", msg)
	}
}

func (cc *CommController) broadcastToClients() {
    //Writing to all client Web Sockets
	for {
		msg := <-cc.broadcast //blocking
		log.Println("Server broadcasted", msg)
		for client := range cc.clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				delete(cc.clients, client)
				client.Close()
			}
		}
	}
}