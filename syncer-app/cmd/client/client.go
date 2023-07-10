package main

import (
	"flag"
	"net/url"

	"github.com/bharadwaja-rao-d/syncing/client"
	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/rs/zerolog"
)

func main() {

	debug := flag.Bool("debug", false, "sets log level to debug")
	flag.Parse()
	args := flag.Args()
	differ := diff.NewDiffer()

	zerolog.SetGlobalLevel(zerolog.WarnLevel)
	if *debug {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}

	integrator := client.NewIntegrator();
	go integrator.Read(differ.FromStd)
	go integrator.Write(differ.ToStd)

	u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/collaborate/" + args[0]}
	c, fmsg := client.NewClient(u, differ)

	go differ.StartDiffer(fmsg)
	c.Start()
}
