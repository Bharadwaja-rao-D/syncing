package main

import (
	"bufio"
	"flag"
	"fmt"
	"net/url"
	"os"

	"github.com/bharadwaja-rao-d/syncing/client"
	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/bharadwaja-rao-d/syncing/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
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

		integrator := NewIntegrator()
		go integrator.read(differ.FromStd)
		go integrator.write(differ.ToStd)

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/collaborate/1"}
		c, fmsg := client.NewClient(u, differ)

		go differ.StartDiffer(fmsg)
		c.Start()
	}

}

type Integrator struct {
	reader *bufio.Reader
	writer *bufio.Writer
	buf    chan string
}

func NewIntegrator() *Integrator {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	buf := make(chan string)

	return &Integrator{reader: reader, writer: writer, buf: buf}
}

// will read from stdin and send to *to chan*
func (i *Integrator) read(to chan string) {
    log.Info().Msg("Started reader")
	for {
		input, _ := i.reader.ReadString('\n')
		to <- input
	}
}

// will read from *from chan* and print to stdout
func (i *Integrator) write(from chan string) {
    log.Info().Msg("Started writer")
    for msg := range from {
        log.Debug().Msgf("Writer: %s\n",msg);
        i.writer.WriteString(msg)
        i.writer.Flush()
    }
}
