package main

import "log"

func main(){
	testingGin()
	//mainActual()
	//testDB()
}

//run differently depending if tables exist
func mainActual(){
	var db Database
	db.OpenConnection()
	db.CreateTables()
	db.PostUser("Georges", "p1")
	db.PostUser("Elias", "p2")
	db.CloseConnection()
	var cc CommController
	cc.Start()
}

func testDB() {
	log.Println("--- Starting testDB Routine ---")
	var db Database
	db.OpenConnection()
	defer db.CloseConnection()
	
	log.Println("Deleting tables")
	db.DeleteTables()
	log.Println("Deleted tables")

	log.Println("Creating tables")
	db.CreateTables()
	log.Println("Created tables")

	db.PostUser("Georges", "p1")
	db.PostUser("Elias", "p2")

	message_id, timestamp:=db.PostMessage(1,"raw","text","morse")
	log.Println("message_id", message_id, "timestamp:",timestamp)

	message_id, timestamp=db.PostMessage(2,"raw","text","morse")
	log.Println("message_id", message_id, "timestamp:",timestamp)

	log.Println("--- Finished testDB Routine ---")
}