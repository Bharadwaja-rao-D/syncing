package main

import (
	"net/url"
	"os"

	"github.com/bharadwaja-rao-d/syncing/client"
	"github.com/bharadwaja-rao-d/syncing/server"
)

func main() {
	args := os.Args

	if args[1] == "server" {
		server.Server()
	} else if args[1] == "client" {
		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/collaborate"}
		c := client.NewClient(u)
        c.FromServer();
	}
}
