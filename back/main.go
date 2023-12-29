package main

import (
	"errors"
	"net/http"

	"github.com/GrandOichii/messager-app/back/router"
	"github.com/gorilla/websocket"
)

func launchRouter() {
	r := router.CreateRouter()

	r.Engine.Run()

}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func ws(w http.ResponseWriter, r *http.Request) {
	// the connection
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		panic(errors.New("failed websocket connection"))
	}

	for {
		// message reading
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			// consider disconnected
			panic(errors.New("failed reading new message from websocket"))
		}

		// message writing
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			panic(errors.New("failed to write message to websocket"))
		}
	}
}

func testWebSocket() {
	http.HandleFunc("/ws", ws)
	http.ListenAndServe("localhost:8080", nil)
}

func main() {
	launchRouter()
	// testWebSocket()
}
