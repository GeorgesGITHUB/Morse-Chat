package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true //Any origin is allowed
	},
}
var clients = make(map[*websocket.Conn]bool) // key points to a Web Socket
var broadcast = make(chan Message)

func handleConnections(w http.ResponseWriter, r *http.Request) {
	//Initial Client Web Socket Connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("WebSocket upgrade error:", err)
		return
	}
	defer conn.Close()

    // Registering client's connection as true
	clients[conn] = true
	log.Println("Client connected")

	//Reading client's message from the Web Socket
	for {
		var msg Message
		err := conn.ReadJSON(&msg) //blocking
		if err != nil {
			log.Println("Error reading JSON:", err)
			delete(clients, conn)
			return
		}
		msg.fillMissingUsingRaw()
		broadcast <- msg //blocking
		log.Println("Server received", msg)
	}
}

func handleMessages() {
    //Writing to all client's Web Sockets
	for {
		msg := <-broadcast //blocking
		log.Println("Server sent", msg)
		for client := range clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				delete(clients, client)
				client.Close()
			}
		}
	}
}

func main() {
	//Invokes Web Socket connection
	http.HandleFunc("/ws", handleConnections) //forever looping
	go handleMessages()                       	 //forever looping
	log.Println("WebSocket server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil)) //loops until error
}
