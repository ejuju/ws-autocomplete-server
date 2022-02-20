package ws

import (
	"fmt"
	"net/http"
	"strings"

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
		_, ok := err.(websocket.HandshakeError)
		if ok {
			fmt.Println("WS handshake error: ", err)
			return
		}
		fmt.Println("Upgrade error: ", err)
		return
	}
	defer conn.Close()

	// listen for incoming client messages
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("an error occured when handling ws message", err)
			return
		}

		str := string(msg)
		results := suggest.End(str, 30)

		// append input before result string
		var withPrefix []string
		for _, res := range results {
			withPrefix = append(withPrefix, str+res)
		}

		fmtStr := strings.Join(withPrefix, "\n")

		conn.WriteMessage(websocket.TextMessage, []byte(fmtStr))
	}
}
