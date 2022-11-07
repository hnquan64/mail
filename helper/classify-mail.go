package helper

import (
	"strings"
)

const format = "be.earning"

type Connection struct {
	Host     string   `json:"host"`
	Port     string   `json:"port"`
	Classify string   `json:"classify"`
	Receiver []string `json:"receiver"`
}

func ClassifyEmail(receiver string) *Connection {
	if strings.Contains(receiver, format) {
		return &Connection{
			Host:     "127.0.0.1",
			Port:     "1025",
			Classify: "System",
			Receiver: []string{receiver},
		}
	}
	return &Connection{
		Host:     "smtp.gmail.com",
		Port:     "587",
		Classify: "Outside",
		Receiver: []string{receiver},
	}
}
