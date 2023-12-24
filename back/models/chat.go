package models

type Chat struct {
	ID                 string     `json:"_id"`
	ParticipantHandles []string   `json:"participants"`
	Messages           []*Message `json:"messages"`
}

func (c *Chat) HasParticipant(handle string) bool {
	for _, p := range c.ParticipantHandles {
		if p == handle {
			return true
		}
	}
	return false
}

type CreateChat struct {
	WithHandle string `json:"with"`
}
