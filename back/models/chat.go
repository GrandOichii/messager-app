package models

type Chat struct {
	ID                 string     `json:"id"`
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
	// TODO remove
	ByHandle string `json:"by"`

	WithHandle string `json:"with"`
}
