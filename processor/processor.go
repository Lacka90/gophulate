package processor

import (
	"bytes"
	"fmt"

	"github.com/Lacka90/gophulate/interfaces"
)

// Processor struct
type Processor struct {
	compute func(int) int
}

// Comp func
func (p *Processor) Comp(comp func(int) int) {
	p.compute = comp
}

// Process - process parsed records
func (p *Processor) Process(clients []*interfaces.Record, mode string) string {
	var stream = make(chan *interfaces.Record)
	var done = make(chan string)

	go consume(p, len(clients), stream, done)

	produce(stream, clients, mode)

	message := <-done

	return message
}

func consume(p *Processor, length int, stream chan *interfaces.Record, done chan string) {
	count := 0
	var messageBuffer bytes.Buffer
	for {
		client := <-stream
		count++
		if p.compute == nil {
			panic("Missing compute propery")
		}
		result := p.compute(client.Iteration)
		messageBuffer.WriteString(fmt.Sprintf("ID: %s, Iteration: %d - Result: %d %d\n", client.ID, client.Iteration, result, count))

		if length == count {
			done <- messageBuffer.String()
		}
	}
}

func produce(stream chan *interfaces.Record, clients []*interfaces.Record, mode string) {
	for i := range clients {
		val := clients[i]
		if mode == "parallel" {
			go func() {
				stream <- val
			}()
		} else {
			stream <- val
		}
	}
}
