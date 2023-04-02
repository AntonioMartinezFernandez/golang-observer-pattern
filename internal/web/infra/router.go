package infra

import (
	"github.com/gorilla/mux"

	mb_domain "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/domain"
	web_application_handlers "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/web/application/handlers"
)

func CreateMuxRouter(mb mb_domain.MessageBroker) *mux.Router {
	mux := mux.NewRouter()

	mux.Handle("/healthcheck", web_application_handlers.NewHealthcheck()).Methods("GET")
	mux.Handle("/message", web_application_handlers.NewMessage(mb)).Methods("POST")
	mux.Handle("/command", web_application_handlers.NewCommand(mb)).Methods("POST")

	return mux
}
