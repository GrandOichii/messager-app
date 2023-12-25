package models

import "errors"

type Message struct {
	OwnerHandle string `json:"uhandle" bson:"uhandle"`
	Text        string `json:"text" bson:"text"`
}

func (m *Message) CheckValid() error {
	if len(m.Text) == 0 {
		return errors.New("can't send text message with no text")
	}

	return nil
}

type PostMessage struct {
	ChatID string `json:"chat_id"`
	Text   string `json:"text"`
}
