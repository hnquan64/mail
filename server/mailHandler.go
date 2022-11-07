package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/emersion/go-sasl"
	"github.com/emersion/go-smtp"
	"gitlab.com/meta-node/mail/core/entities"
	"gitlab.com/meta-node/mail/helper"
)

type MailHandler struct {
	Flag  chan error
	Email *entities.Email
	Auth  sasl.Client
}

func (mh *MailHandler) SendmailHandler() {
	for _, receiver := range mh.Email.To {
		conn := helper.ClassifyEmail(receiver)
		log.Println("Connection classify is ", conn.Classify)
		switch conn.Classify {
		case "System":
			{
				go mh.sendSystemEmail(conn)
			}
		case "Outside":
			{
				go mh.sendOutsideEmail(conn)
			}
		}
	}
}

func (mh *MailHandler) sendSystemEmail(conn *helper.Connection) {
	address := fmt.Sprintf("%v:%v", conn.Host, conn.Port)
	content := strings.NewReader(mh.Email.Content)
	if err := smtp.SendMail(address, mh.Auth, mh.Email.From, conn.Receiver, content); err != nil {
		log.Println(err)
		return
	}
	mh.Email.Status = SENT
	dbConn.DB.Save(&mh.Email)
}

func (mh *MailHandler) sendOutsideEmail(conn *helper.Connection) {
	address := fmt.Sprintf("%v:%v", conn.Host, conn.Port)
	content := strings.NewReader(mh.Email.Content)
	if err := smtp.SendMail(address, mh.Auth, mh.Email.From, conn.Receiver, content); err != nil {
		log.Println(err)
		return
	}
	mh.Email.Status = SENT
	dbConn.DB.Save(&mh.Email)
}

// func (mh MailHandler) ReceivemailHandler() {}
