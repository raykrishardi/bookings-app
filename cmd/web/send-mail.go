package main

import (
	"log"
	"time"

	"github.com/tsawler/bookings-app/internal/models"
	mail "github.com/xhit/go-simple-mail/v2"
)

func listenForMail() {
	go func() {
		for {
			m := <-app.MailChan
			sendMsg(m)
		}
	}()
}

func sendMsg(m models.MailData) {
	server := mail.NewSMTPClient()
	server.Host = "localhost"
	server.Port = 1025
	server.KeepAlive = false // Not to keep connection to the mail server active all the time, only make a connection when required
	server.ConnectTimeout = 10 * time.Second
	server.SendTimeout = 10 * time.Second

	// Example of additional arguments for prod
	// server.Username = ""
	// server.Password = ""
	// server.Encryption = ""

	client, err := server.Connect()
	if err != nil {
		errorLog.Println(err)
	}

	email := mail.NewMSG()
	email.SetFrom(m.From).AddTo(m.To).SetSubject(m.Subject)
	email.SetBody(mail.TextHTML, m.Content)

	err = email.Send(client)
	if err != nil {
		errorLog.Println(err)
	}

	log.Println("Email sent!")
}
