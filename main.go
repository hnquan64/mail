package main

import (
	"gitlab.com/meta-node/mail/config"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/routers"
	"gitlab.com/meta-node/mail/server"
)

func main() {
	finish := make(chan bool)
	go runHttpsServer()
	go server.MailServer()

	<-finish
}

func runHttpsServer() {
	dbConn := database.InstanceDB()
	database.Migration(dbConn)
	r := routers.InitRouter()
	port := config.GetConfig().Server.HTTP.Port
	r.Run(port)
}
