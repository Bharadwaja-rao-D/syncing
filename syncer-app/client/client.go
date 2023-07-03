package client

import (
	"net/url"

	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/bharadwaja-rao-d/syncing/protocol"
	"github.com/gorilla/websocket"

	"github.com/rs/zerolog/log"
)

type Client struct {
    uname string
	conn   *websocket.Conn
	differ *diff.Differ
}

func NewClient(url url.URL, differ *diff.Differ) (*Client, string) {
	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		log.Fatal().Err(err)
	}

    var fmsg protocol.CSMessage
    _, msg, _ := conn.ReadMessage();
    log.Debug().Msg(string(msg))
    err = conn.ReadJSON(fmsg)
	if err != nil {
		log.Fatal().Err(err)
	}
    log.Debug().Msgf("NewClient %s", fmsg.Content)
	return &Client{conn: conn, differ: differ}, fmsg.Content
}

//recvs messages from server and sends to *FromClient chan*
func (c *Client) fromServer() {
	d := c.differ
	conn := c.conn
    var msg protocol.CSMessage
	for {
		err := conn.ReadJSON(msg)
		if err != nil {
		log.Fatal().Err(err)
			break
		}
        log.Debug().Msgf("Client: %s\n", msg.Content)
		d.FromClient <- diff.EditScript(msg.Content)
	}
}

//sends messages from *ToClient chan* to server
func (c *Client) toServer() {
	d := c.differ
	conn := c.conn
	for msg := range d.ToClient {
        conn.WriteJSON(protocol.CSMessage[diff.EditScript]{Mtype: protocol.Update, Content: msg, From: conn.LocalAddr().String()})
	}
}

func (c *Client) Start() {
	log.Debug().Msg("Started Client")
    go c.toServer()
    c.fromServer()
}
