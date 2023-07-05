package main

import (
	"flag"
	"fmt"
	"net/url"

	"github.com/bharadwaja-rao-d/syncing/client"
	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/bharadwaja-rao-d/syncing/server"
	"github.com/rs/zerolog"
)

func main() {

	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	args := flag.Args()

	// Default level for this example is info, unless debug flag is present
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if *debug {
		fmt.Println("Debug set true")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	if args[0] == "server" {
		server.Server()
	} else if args[0] == "client" {

		differ := diff.NewDiffer()

		integrator := client.NewIntegrator()
		go integrator.Read(differ.FromStd)
		go integrator.Write(differ.ToStd)

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/collaborate/" + args[1]}
		c, fmsg := client.NewClient(u, differ)

		go differ.StartDiffer(fmsg)
		c.Start()
	}

}
