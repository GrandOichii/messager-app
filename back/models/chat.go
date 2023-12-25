package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Chat struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	ParticipantHandles []string           `json:"participants" bson:"participants"`
	Messages           []*Message         `json:"messages" bson:"messages"`
}

func (c *Chat) HasParticipant(handle string) bool {
	for _, p := range c.ParticipantHandles {
		if p == handle {
			return true
		}
	}
	return false
}

func (c *Chat) ToGetChat() *GetChat {
	return &GetChat{
		ID:                 c.ID.Hex(),
		ParticipantHandles: c.ParticipantHandles,
		Messages:           c.Messages,
	}
}

type GetChat struct {
	ID                 string     `json:"id"`
	ParticipantHandles []string   `json:"participants"`
	Messages           []*Message `json:"messages"`
}

type CreateChat struct {
	WithHandle string `json:"with"`
}
