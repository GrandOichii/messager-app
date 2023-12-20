package models

type Message struct {
	// TODO use user id?
	OwnerHandle string `json:"uhandle"`
	Text        string `json:"text"`
}
