package models

type Message struct {
	// TODO use user id?
	OwnerHandle string `json:"uhandle"`
	Text        string `json:"text"`
}

type PostMessage struct {
	ChatID      string `json:"chat_id"`
	Text        string `json:"text"`
	OwnerHandle string `json:"uhandle"`
}
