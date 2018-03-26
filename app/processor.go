package app

import (
	"bytes"
	"fmt"
)

func fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}

func consume(length int, stream chan *Record, done chan string) {
	count := 0
	var messageBuffer bytes.Buffer
	for {
		client := <-stream
		count++
		result := fibonacci(client.Iteration)
		messageBuffer.WriteString(fmt.Sprintf("ID: %s, Iteration: %d - Result: %d %d\n", client.ID, client.Iteration, result, count))

		if length == count {
			done <- messageBuffer.String()
		}
	}
}

func produce(stream chan *Record, clients []*Record, mode string) {
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

// Process - process parsed records
func Process(clients []*Record, mode string) string {
	var stream = make(chan *Record)
	var done = make(chan string)

	go consume(len(clients), stream, done)

	produce(stream, clients, mode)

	message := <-done

	return message
}
