package main

import (
	"log"
	"net/http"
	"github.com/gorilla/websocket"
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

func (cc *CommController) listenForClient(conn *websocket.Conn){
	//Reading from client's Web Socket
	for {
		var msg Message
		err := conn.ReadJSON(&msg) //blocking

		if err != nil {
			log.Println("Failed reading JSON from Web Socket:", err)
			delete(cc.clients, conn)
			log.Println("UnRegistering client's WebSocket Connection:")
			return
		}

		msg.fillMissingUsingRaw()

		// TODO
			// Insert new message to table

		cc.broadcast <- msg //blocking
		log.Println("HTTP Server received", msg)
	}
}

func (cc *CommController) broadcastToClients() {
    //Writing to all client Web Sockets
	for {
		msg := <-cc.broadcast //blocking
		log.Println("Server sent", msg)
		for client := range cc.clients {
			if err := client.WriteJSON(msg); err != nil {
				log.Println("Error writing JSON:", err)
				delete(cc.clients, client)
				client.Close()
			}
		}
	}
}