package client

import (
	"log"
	"net/url"

	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	differ *diff.Differ
}

func NewClient(url url.URL) *Client {
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal(err)
	}
    conn.WriteMessage(websocket.TextMessage,[]byte("Hello Server"));
	return &Client{Conn: conn, differ: new(diff.Differ)}
}

func (c *Client) Watch(file_path string) {
}

// These two functions should run concurrently
func (c *Client) ToServer() {}

func (c *Client) FromServer() {}
