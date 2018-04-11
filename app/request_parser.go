package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/Lacka90/gophulate/computations"
)

// GetProcessMode - get process mode from request
func GetProcessMode(r *http.Request) (string, error) {
	mode := r.FormValue("mode")

	switch mode {
	case "serial":
	case "parallel":
		return mode, nil
	default:
		return "", fmt.Errorf("Process mode not supported: %s", mode)
	}
	return "", nil
}

// GetProcessor - get processor type from request
func GetProcessor(r *http.Request) (func(int) string, error) {
	processor := r.FormValue("processor")

	switch processor {
	case "fibonacci":
		return computations.Fibonacci, nil
	case "isprime":
		return computations.IsPrime, nil
	default:
		return nil, fmt.Errorf("Processor not supported: %s", processor)
	}
	return nil, nil
}

// ReadFileContent - get file content from request
func ReadFileContent(r *http.Request) string {
	var Buf bytes.Buffer

	file, _, err := r.FormFile("file")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	io.Copy(&Buf, file)

	contents := Buf.String()

	Buf.Reset()

	return contents
}
