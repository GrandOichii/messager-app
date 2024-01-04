package connection

import (
	"encoding/json"
	"errors"
	"net/http"
	"slices"

	"github.com/GrandOichii/messager-app/back/models"
	"github.com/GrandOichii/messager-app/back/services"
	"github.com/gorilla/websocket"
	"gopkg.in/dgrijalva/jwt-go.v3"
)

type MessageHub interface {
	// Run()
	Register(chatID string, w http.ResponseWriter, r *http.Request) error
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

	Services *services.Services

	upgrader websocket.Upgrader
	chatMap  map[string][]*Client

	register chan *QueuedRegister
	notify   chan *QueuedNotify
}

func NewNotifyHub(services *services.Services) *NotifyHub {
	result := &NotifyHub{
		Services: services,
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
			// * there is no check for multiple clients with the same handle, i think this should stay
			nh.chatMap[qr.ChatID] = append(nh.chatMap[qr.ChatID], qr.c)
		case m := <-nh.notify:
			clients, ok := nh.chatMap[m.ChatID]
			if !ok {
				// TODO
				panic(errors.New("requested to notify users about chat message, but no users found"))
			}

			// create a new list with all the clients that received the message successfully, then assign this new list as the client list
			newClients := []*Client{}
			for _, client := range clients {
				data, err := json.Marshal(m.Message)
				if err != nil {
					// * shouldn't throw any errors, as the marshal method breaks only when the data is cyclical
					panic(err)
				}
				err = client.Connection.WriteMessage(client.MessageType, data)
				if err == nil {
					newClients = append(newClients, client)
				}
			}
			nh.chatMap[m.ChatID] = newClients
		}
	}
}

func (nh *NotifyHub) Register(chatID string, w http.ResponseWriter, r *http.Request) error {

	conn, err := nh.upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(err)
	}

	// read jwt token
	mT, p, err := conn.ReadMessage()
	if err != nil {
		panic(err)
	}

	token := string(p)
	claims := jwt.MapClaims{}
	jwt.ParseWithClaims(token, claims, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret key"), nil
	})
	handle := claims["handle"].(string)

	_, err = nh.Services.UserServicer.ByHandle(handle)
	if err != nil {
		return err
	}

	chat, err := nh.Services.ChatServicer.ByID(chatID)
	if err != nil {
		return err
	}

	contains := slices.Contains(chat.ParticipantHandles, handle)
	if !contains {
		return errors.New("requested to listen for chat messages when not being a participant")
	}

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
