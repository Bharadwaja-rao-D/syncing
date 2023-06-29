package main

import (
	"bufio"
	"log"
	"net/url"
	"os"

	"github.com/bharadwaja-rao-d/syncing/client"
	"github.com/bharadwaja-rao-d/syncing/diff"
	"github.com/bharadwaja-rao-d/syncing/server"
)

func main() {
	args := os.Args

	if args[1] == "server" {
		server.Server()
	} else if args[1] == "client" {

		differ := diff.NewDiffer()
		go differ.StartDiffer()

		integrator := NewIntegrator()
		go integrator.read(differ.FromStd)
		go integrator.write(differ.ToStd)

		u := url.URL{Scheme: "ws", Host: "localhost:8080", Path: "/collaborate"}
		c := client.NewClient(u, differ)
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
	log.Println("Starting reader ...")
	for {
		input, _ := i.reader.ReadString('\n')
		to <- input
		if input == "exit" {
			break
		}
	}
}

// will read from *from chan* and print to stdout
func (i *Integrator) write(from chan string) {
	log.Println("Starting writer ...")
	for {
		msg := <-from
		if msg == "exit" {
			break
		}
		i.writer.WriteString(msg)
		i.writer.Flush()
	}
}
