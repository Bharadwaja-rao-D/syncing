package server

import (
	"net/http"
	"sync"

	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/bharadwaja-rao-d/syncing/protocol"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

type SharedState struct {
	mu    sync.Mutex
	Rooms map[string]*Room
}

func (s *SharedState) AddRoom(room *Room) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.Rooms[room.room_id] = room
	log.Info().Msgf("INFO: Added room with id %s\n", room.room_id)
}

func (s *SharedState) GetRoom(room_id string) *Room {
	//no need of lock
	s.mu.Lock()
	defer s.mu.Unlock()

	return s.Rooms[room_id]
}

func start_collaboration(w http.ResponseWriter, r *http.Request, s *SharedState) {
	//create a new room and return its id
	room := NewRoom()
	s.AddRoom(room)
	w.Write([]byte(room.room_id))

}

func collaborate(w http.ResponseWriter, r *http.Request, s *SharedState) {

	room_id := r.URL.Path[len("/collaborate/"):]
	room := s.GetRoom(room_id)

	upgrader := new(websocket.Upgrader)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

	room.conns = append(room.conns, conn)
	log.Info().Msgf("%s joined room with id %s\n", conn.RemoteAddr(), room.room_id)

	defer conn.Close()

	//conn.WriteMessage(websocket.TextMessage, []byte(room.last_file))
    fmsg := protocol.CSMessage[string]{Mtype: protocol.HandShake, From: "server", Content: room.lastest_content}
    log.Debug().Msgf(fmsg.Content)
    conn.WriteJSON(fmsg)

	for {

        var msg protocol.CSMessage[diff.EditScript];
        _ = conn.ReadJSON(msg)
		if  len(msg.Content) == 0 {
			log.Info().Msgf("Closing connection: %s\n", conn.RemoteAddr())
			break
		}

        //updating the latest file u have in server
		room.lastest_content = room.differ.FromDiff(room.lastest_content, msg.Content)

		//broadcasting to all clients of the group
		for _, clients := range room.conns {
            err := clients.WriteJSON(protocol.CSMessage[diff.EditScript]{Mtype: protocol.Update, From: msg.From, Content: msg.Content})
			if err != nil {
				log.Fatal().Err(err)
			}
			log.Printf("To client %s: %s\n", conn.RemoteAddr(), msg.Content)
		}
	}
}

func Server() {

	state := &SharedState{Rooms: make(map[string]*Room)}

	http.HandleFunc("/collaborate", func(w http.ResponseWriter, r *http.Request) {
		start_collaboration(w, r, state)
	})
	http.HandleFunc("/collaborate/", func(w http.ResponseWriter, r *http.Request) {
		collaborate(w, r, state)
	})
	http.ListenAndServe(":8080", nil)
}
