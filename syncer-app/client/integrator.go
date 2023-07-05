package client

import (
	"bufio"
	"os"

	"github.com/rs/zerolog/log"
)

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

// will Read from stdin and send to *to chan*
func (i *Integrator) Read(to chan string) {
    log.Info().Msg("Started reader")
	for {
		input, _ := i.reader.ReadString('\n')
		to <- input
	}
}

// will read from *from chan* and print to stdout
func (i *Integrator) Write(from chan string) {
    log.Info().Msg("Started writer")
    for msg := range from {
        log.Debug().Msgf("Writer: %s\n",msg);
        i.writer.WriteString(msg)
        i.writer.Flush()
    }
}
