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
	var db AWS_RDS
	db.openConnection()
	defer db.closeConnection()
	db.createUsersTable()
	db.createMessagesTable()
	a:= db.addUser("Georges")
	b:= db.addUser("Quandale")
	db.addMessage(a,"raw1", "text1", "morse1")
	db.addMessage(a,"raw2", "text2", "morse2")
	db.addMessage(b,"raw1", "text1", "morse1")
	db.addMessage(b,"raw2", "text2", "morse2")
	//db.deleteAllTables()

}