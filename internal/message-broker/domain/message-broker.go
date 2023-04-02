package domain

type MessageBroker interface {
	ProcessEvent()
	PublishEvent(event Event)
}
