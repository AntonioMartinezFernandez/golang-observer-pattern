package domain

type Event struct {
	Type    string `json:"type"`
	Content string `json:"content"`
}
