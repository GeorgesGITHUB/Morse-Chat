package main

import "log"

func main(){ 
	// mainActual()
	//testDB()
}

func mainActual(){
	// log.Println("--- Setting up a fresh DB w/ Schemas ---")
	// var db AWS_RDS
	// db.openConnection()
	// db.deleteAllTables()
	// db.createUsersTable()
	// db.createMessagesTable()
	// db.addUser(613,"Georges")
	// db.addUser(961,"Elias")
	// db.addUser(627,"John")
	// db.closeConnection()

	log.Println("--- Starting main ---")
	var cc CommController
	cc.Start()
}

func testDB() {
	// log.Println("--- Starting testDB Routine ---")
	// var db AWS_RDS
	// db.openConnection()
	// defer db.closeConnection()
	// db.deleteAllTables()
	// db.createUsersTable()
	// db.createMessagesTable()
	// db.addUser(961,"Georges")
	// db.addUser(627,"John")
	// message_id, timestamp:=db.addMessage(961,"raw","text","morse")
	// log.Println("message_id", message_id, "timestamp:",timestamp)
	// message_id, timestamp=db.addMessage(627,"raw","text","morse")
	// log.Println("message_id", message_id, "timestamp:",timestamp)
}