package main

import (
	"log"
	"os"

	"github.com/AntonioMartinezFernandez/golang-observer-pattern/config"
	mb_application "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/application"
	mb_domain "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/domain"

	web_infra "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/web/infra"
)

func main() {
	// Start Log
	log.SetOutput(os.Stdout)
	log.Println("[APP] starting application...")

	// Create Main Channel (with buffer)
	mainChannel := make(chan mb_domain.Event, 10)

	// Stop Gracefully
	defer func() {
		log.Println("[APP] stopping application...")
		close(mainChannel)
	}()

	// Load Environment Variables
	config.LoadEnvVariables()
	HTTP_PORT := os.Getenv("HTTP_PORT")

	// Start Message Broker
	mb := mb_application.NewMessageBroker(mainChannel)
	go mb.ProcessEvent()

	// Start Http Server
	web_infra.Start(HTTP_PORT, mb)
}
