package infra

import (
	"log"
	"net/http"
	"time"

	mb_domain "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/domain"
)

func Start(port string, mb mb_domain.MessageBroker) *http.Server {
	// Mux Router
	mux := CreateMuxRouter(mb)

	// Create Server
	server := &http.Server{
		Addr:        "0.0.0.0:" + port,
		Handler:     mux,
		ReadTimeout: time.Second * 5,
		IdleTimeout: time.Second * 5,
	}

	// Start Server
	log.Println("[HTTP] starting http server on localhost:" + port)
	log.Fatal(server.ListenAndServe())

	return server
}
