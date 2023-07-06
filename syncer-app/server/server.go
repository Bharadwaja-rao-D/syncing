package server

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
    "github.com/rs/cors"
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

    log.Debug().Msg(r.Header.Get("Upgrade"));

	upgrader := new(websocket.Upgrader)
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal().Err(err)
	}

    log.Debug().Msgf("%p, %p", room.conns, conn);

	room.conns = append(room.conns, conn)
	log.Info().Msgf("%s joined room with id %s\n", conn.RemoteAddr(), room.room_id)

	defer conn.Close()

	conn.WriteMessage(websocket.TextMessage, []byte(room.last_file))

	for {

		mtype, msg, _ := conn.ReadMessage()
		if mtype == websocket.CloseMessage || len(msg) == 0 {
			log.Info().Msgf("Closing connection: %s\n", conn.RemoteAddr())
			break
		}

		room.last_file = room.differ.FromDiff(room.last_file, string(msg))

		//broadcasting to all clients of the group
		for _, clients := range room.conns {
			err := clients.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Fatal().Err(err)
			}
			log.Printf("To client %s: %s\n", conn.RemoteAddr(), msg)
		}
	}
}


func Server() {
    mux := http.NewServeMux()

	state := &SharedState{Rooms: make(map[string]*Room)}

	mux.HandleFunc("/collaborate", func(w http.ResponseWriter, r *http.Request) {
		start_collaboration(w, r, state)
	})

	mux.HandleFunc("/collaborate/", func(w http.ResponseWriter, r *http.Request) {
		collaborate(w, r, state)
	})

    // cors.Default() setup the middleware with default options being
    // all origins accepted with simple methods (GET, POST). See
    // documentation below for more options.
    handler := cors.Default().Handler(mux)
    http.ListenAndServe(":8080", handler)
}
