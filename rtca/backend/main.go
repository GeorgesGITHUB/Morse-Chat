package main

import (
	"log"
	"net/http"
)

func main(){ mainTest() }

func mainActual(){
	//Invokes Web Socket connection
	http.HandleFunc("/ws", handleConnections) //forever looping
	go handleMessages()                       	 //forever looping
	log.Println("WebSocket server started at ws://localhost:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil)) //loops until error
}

func mainTest() {
	testRDS() //testing db
}