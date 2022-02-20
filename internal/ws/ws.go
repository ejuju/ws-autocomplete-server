package ws

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ejuju/ws-autocomplete-server/internal/suggest"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(*http.Request) bool { return true }, // todo
}

// Serve upgrades the http request to use the websocket protocol and listens for incoming messages from the client
func Serve(w http.ResponseWriter, r *http.Request) {

	// Upgrade request to use websocket protocol
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println("WS handshake error: ", err)
			return
		}
		log.Println("Upgrade error: ", err)
		return
	}
	defer conn.Close()

	// listen for incoming client messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Println("an error occured when handling ws message", err)
			return
		}

		if len(msg) == 0 {
			continue
		}

		str := string(msg)
		fmt.Println(str)

		results, err := suggest.End(str, 30)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Println(results)
	}
}
