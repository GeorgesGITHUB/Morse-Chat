package main

import (
	"log"
	"github.com/gin-gonic/gin"
)

func main(){
	var db Database
	db.OpenConnection()
	db.DeleteTables()
	db.CreateTables()
	db.CloseConnection()

	router := gin.Default()
	
	enableCORS(router)
	registerAPItoEndpoint(router)
	registerWStoEndpoint(router)
	
	router.Run(":8080")
}

// use in a test suite later
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