package client

import (
	"log"
	"net/url"

	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/gorilla/websocket"
)

type Client struct {
	conn   *websocket.Conn
	differ *diff.Differ
}

func NewClient(url url.URL, differ *diff.Differ) *Client {
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
	return &Client{conn: conn, differ: differ}
}

//recvs messages from server and sends to *FromClient chan*
func (c *Client) fromServer() {
	d := c.differ
	conn := c.conn
	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
			break
		}
        log.Printf("DEBUG: fromServer %s\n", msg)
		d.FromClient <- string(msg)
	}
}

//sends messages from *ToClient chan* to server
func (c *Client) toServer() {
	d := c.differ
	conn := c.conn
	for msg := range d.ToClient {
		conn.WriteMessage(websocket.TextMessage, []byte(msg))
	}
}

func (c *Client) Start() {
	log.Println("Started Client...")
    go c.toServer()
    c.fromServer()
}
