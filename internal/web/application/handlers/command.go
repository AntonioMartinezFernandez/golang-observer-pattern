package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	mb_domain "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/domain"
)

type Command struct {
	messageBroker mb_domain.MessageBroker
}

func NewCommand(mb mb_domain.MessageBroker) *Command {
	return &Command{
		messageBroker: mb,
	}
}

func (c Command) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Decode Request Body (JSON) As Data Struct
	var body interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Http Response
	w.WriteHeader(http.StatusCreated)

	// Marshall Body Struct And Convert To String
	bodyAsStruct, _ := json.Marshal(body)

	bodyAsString := string(bodyAsStruct)
	log.Println("[HTTP] received command body: ", bodyAsString)

	// Event to Publish
	event := mb_domain.Event{
		Type:    "command",
		Content: bodyAsString,
	}

	// Publish Event To Channel
	go c.messageBroker.PublishEvent(event)
}
