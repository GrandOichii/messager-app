package models

type Message struct {
	OwnerHandle string `json:"uhandle"`
	Text        string `json:"text"`
}

type PostMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
