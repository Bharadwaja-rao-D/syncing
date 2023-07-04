package server

import (
	"strconv"

	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/gorilla/websocket"
)

var id_gen = 0

// TODO: Might be needing mutex
type Room struct {
	room_id   string
	conns     []*websocket.Conn
	last_file string
	differ    diff.Differ
}

func NewRoom() *Room {
	id_gen += 1
	return &Room{room_id: strconv.FormatInt(int64(id_gen), 10), conns: make([]*websocket.Conn, 0), differ: *diff.NewDiffer()}
}

func (r *Room) Start() {
}
