package client

import (
	"bufio"
	"os"

	"github.com/bharadwaja-rao-d/syncing/message"
	"github.com/rs/zerolog/log"
)

type Integrator struct {
	reader *bufio.Reader
	writer *bufio.Writer
}

func NewIntegrator() *Integrator {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)

	return &Integrator{reader: reader, writer: writer}
}

// will Read from stdin and send to *to chan*
func (i *Integrator) Read(to chan string) {
	for {
		input, _ := i.reader.ReadString('\n')
        std_msg := message.Deserialize(input);

        var server_text string
        for txt := range std_msg.Text{
            server_text += std_msg.Text[txt]
        }
		to <- server_text
	}
}

// will read from *from chan* and print to stdout
func (i *Integrator) Write(from chan string) {
    for server_text := range from {
        log.Debug().Msgf("Writer: %s",server_text);
        i.writer.WriteString(server_text)
        i.writer.Flush()
    }
}
