package connection

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/gorilla/websocket"
)

type MessageHub interface {
	// Run()
	Register(handle string, chatID string, w http.ResponseWriter, r *http.Request) error
	Notify(chatID string, message *models.Message)
}

// === Implementation ===

type QueuedNotify struct {
	ChatID  string
	Message *models.Message
}

type QueuedRegister struct {
	c      *Client
	ChatID string
}

type Client struct {
	Handle      string
	Connection  *websocket.Conn
	MessageType int
}

type NotifyHub struct {
	// MessageHub

	upgrader websocket.Upgrader
	chatMap  map[string][]*Client

	register chan *QueuedRegister
	notify   chan *QueuedNotify
}

func NewNotifyHub() *NotifyHub {
	result := &NotifyHub{
		upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		chatMap:  map[string][]*Client{},
		register: make(chan *QueuedRegister),
		notify:   make(chan *QueuedNotify),
	}
	go result.Run()
	return result
}

func (nh *NotifyHub) Run() {
	for {
		select {
		case qr := <-nh.register:
			_, ok := nh.chatMap[qr.ChatID]
			if !ok {
				nh.chatMap[qr.ChatID] = []*Client{}
			}
			// TODO check if already exists
			nh.chatMap[qr.ChatID] = append(nh.chatMap[qr.ChatID], qr.c)
		case m := <-nh.notify:
			clients, ok := nh.chatMap[m.ChatID]
			if !ok {
				// TODO
				panic(errors.New("requested to notify users about chat message, but no users found"))
			}
			// TODO easier to just remove?
			newClients := []*Client{}
			for _, client := range clients {
				data, err := json.Marshal(m.Message)
				if err != nil {
					// TODO
					panic(err)
				}
				err = client.Connection.WriteMessage(client.MessageType, data)
				if err == nil {
					// TODO check if there are any errors beside disconnected
					newClients = append(newClients, client)
				}
			}
			nh.chatMap[m.ChatID] = newClients
		}
	}
}

func (nh *NotifyHub) Register(handle string, chatID string, w http.ResponseWriter, r *http.Request) error {

	conn, err := nh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		return err
	}

	// read jwt token
	mT, p, err := conn.ReadMessage()
	if err != nil {
		return err
	}

	token := string(p)

	// TODO check jwt token for validity
	fmt.Printf("token: %v\n", token)

	c := &QueuedRegister{
		ChatID: chatID,
		c: &Client{
			Handle:      handle,
			Connection:  conn,
			MessageType: mT,
		},
	}

	nh.register <- c

	return nil
}

func (nh *NotifyHub) Notify(chatID string, message *models.Message) {
	m := &QueuedNotify{
		ChatID:  chatID,
		Message: message,
	}
	nh.notify <- m
}
