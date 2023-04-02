package application

import (
	"encoding/json"
	"log"

	mb_domain "github.com/AntonioMartinezFernandez/golang-observer-pattern/internal/message-broker/domain"
)

type MessageBroker struct {
	channel chan mb_domain.Event
}

func NewMessageBroker(channel chan mb_domain.Event) *MessageBroker {
	return &MessageBroker{
		channel: channel,
	}
}

func (mb *MessageBroker) ProcessEvent() {
	for event := range mb.channel {
		eventAsJson, err := json.Marshal(event)
		if err != nil {
			log.Printf("[Event Logger] error: %v\n", err)
			return
		}
		log.Printf("[Event Logger] received event: %v\n", string(eventAsJson))
	}
	log.Printf("[Event Logger] event closed")
}

func (mb *MessageBroker) PublishEvent(event mb_domain.Event) {
	mb.channel <- event
}
