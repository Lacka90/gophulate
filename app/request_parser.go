package app

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

// GetProcessMode - get process mode from request
func GetProcessMode(r *http.Request) string {
	return r.FormValue("mode")
}

// ValidateProcessMode - validating process mode
func ValidateProcessMode(mode string) (string, error) {
	switch mode {
	case "serial":
	case "parallel":
		return mode, nil
	default:
		return "", fmt.Errorf("Process mode not supported: %s", mode)
	}
	return "", nil
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
