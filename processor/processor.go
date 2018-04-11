package processor

import (
	"bytes"
	"fmt"

	"github.com/Lacka90/gophulate/interfaces"
)

// Processor struct
type Processor struct {
	compute func(int) string
}

// Comp func
func (p *Processor) Comp(comp func(int) string) {
	p.compute = comp
}

// Process - process parsed records
func (p *Processor) Process(clients []*interfaces.Record, mode string) string {
	var done = make(chan string)
	var progress = make(chan string)

	go listen(progress, done, len(clients))

	produce(p, clients, progress, mode)

	message := <-done

	return message
}

func produce(p *Processor, clients []*interfaces.Record, progress chan string, mode string) {
	for i := range clients {
		value := clients[i]
		if mode == "parallel" {
			go consume(p, value, progress)
		} else {
			consume(p, value, progress)
		}
	}
}

func consume(p *Processor, client *interfaces.Record, progress chan string) {
	if p.compute == nil {
		panic("Missing compute propery")
	}

	result := p.compute(client.Iteration)
	message := fmt.Sprintf("ID: %s, Iteration: %d, Result: %v\n", client.ID, client.Iteration, result)

	progress <- message
}

func listen(progress chan string, done chan string, length int) {
	count := 0
	var messageBuffer bytes.Buffer
	for {
		message := <-progress
		messageBuffer.WriteString(message)
		count++

		if count == length {
			done <- messageBuffer.String()
		}
	}
}
