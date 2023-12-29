package connection

import (
	"fmt"
	"net/http"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/gorilla/websocket"
)

type MessageHub interface {
	// Run()
	Register(handle string, chatID string, w http.ResponseWriter, r *http.Request) error
	Notify(handle string, chatID string, message *models.Message)
}

type Client struct {
	Handle string
}

type NotifyHub struct {
	// MessageHub

	upgrader websocket.Upgrader
	register chan *Client
}

func NewNotifyHub() *NotifyHub {
	result := &NotifyHub{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		register: make(chan *Client),
	}
	go result.Run()
	return result
}

func (nh *NotifyHub) Run() {
	for {
		select {
		case client := <-nh.register:
			fmt.Println("Register request from " + client.Handle)
		}
	}
}

func (nh *NotifyHub) Register(handle string, chatID string, w http.ResponseWriter, r *http.Request) error {
	// TODO
	c := &Client{
		Handle: handle,
	}
	nh.register <- c
	return nil
}

func (nh *NotifyHub) Notify(handle string, chatID string, message *models.Message) {
	// TODO
}