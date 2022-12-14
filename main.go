package main

import (
	"os"

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

	port, existed := os.LookupEnv("PORT")
	if !existed {
		port = "3000"
	}
	r.Run(":" + port)
}
