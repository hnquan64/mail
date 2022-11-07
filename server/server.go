package server

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
	"time"

	"github.com/emersion/go-smtp"
	"gitlab.com/meta-node/mail/core/database"
	"gitlab.com/meta-node/mail/core/entities"
)

var dbConn = database.InstanceDB()

const (
	FORMAT = "be.earning"
	SENT   = "Sent"
)

// The Backend implements SMTP server methods.
type Backend struct{}

// Login handles a login command with username and password.
func (bkd *Backend) Login(state *smtp.ConnectionState, email, password string) (smtp.Session, error) {

	// if email company
	if strings.Contains(email, FORMAT) {
		passwordHash := ""
		if err := dbConn.DB.Model(&entities.User{}).Select("password").Where("email = ?", email).Scan(&passwordHash).Error; err != nil {
			log.Fatal("Not found email address")
			return nil, err
		}

		if password != passwordHash {
			return nil, errors.New("password is	incorrect")
		}
	}

	return &Session{}, nil
}

// AnonymousLogin requires clients to authenticate using SMTP AUTH before sending emails
func (bkd *Backend) AnonymousLogin(state *smtp.ConnectionState) (smtp.Session, error) {
	fmt.Println("ẩn danh ................")
	return nil, smtp.ErrAuthRequired
}

// A Session is returned after successful login.
type Session struct {
	Email entities.Email
}

func (s *Session) Mail(from string, opts smtp.MailOptions) error {
	s.Email.From = from
	return nil
}

func (s *Session) Rcpt(to string) error {
	toSystem := []string{to}
	s.Email.To = toSystem
	return nil
}

func (s *Session) Data(r io.Reader) error {
	if b, err := ioutil.ReadAll(r); err != nil {
		return err
	} else {
		s.Email.Content = string(b)
	}
	return nil
}

func (s *Session) Save() error {
	if err := dbConn.DB.Save(&s.Email).Error; err != nil {
		return err
	}
	return nil
}

func (s *Session) Reset() {}

func (s *Session) Logout() error {
	return nil
}

func MailServer() {
	be := &Backend{}

	s := smtp.NewServer(be)

	s.Addr = ":1025"
	s.Domain = "localhost"
	s.ReadTimeout = 10 * time.Second
	s.WriteTimeout = 10 * time.Second
	s.MaxMessageBytes = 1024 * 1024
	s.MaxRecipients = 50
	s.AllowInsecureAuth = true

	log.Println("Starting server at", s.Domain+s.Addr)
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err.Error())
	}
}
