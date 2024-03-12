package video

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type Peer string

type Room struct {
	Peers []Peer
}

var room Room = Room{}

func GetRoomHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		json.NewEncoder(w).Encode(room)
	}
}

func GetEnterInRoomHandler() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		body, err := io.ReadAll(r.Body)
		if err != nil {
			panic(err)
		}

		peer := Peer(body)

		room.Peers = append(room.Peers, peer)

		log.Println(peer)

		json.NewEncoder(w).Encode("O_K")
	}
}

var connects map[*websocket.Conn]bool = make(map[*websocket.Conn]bool)

func GetSignalHandler() http.HandlerFunc {

	upgrader := websocket.Upgrader{}

	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			panic(err)
		}

		connects[conn] = true

		go readPump(conn)

	}
}

func readPump(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("error: %v", err)
			}
			break
		}

		for anotherConn := range connects {

			if conn != anotherConn {
				anotherConn.WriteMessage(websocket.TextMessage, message)
			}
		}
	}

	delete(connects, conn)
}
