package main

func main(){ mainActual() }

func mainActual(){
	var cc CommController
	cc.Start()
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