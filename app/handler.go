package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handle - file upload handler
func Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mode := GetProcessMode(r)
	mode, err := ValidateProcessMode(mode)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	content := ReadFileContent(r)
	fmt.Fprintf(w, "Mode: %s\nContent: %s", mode, content)
	return
}
