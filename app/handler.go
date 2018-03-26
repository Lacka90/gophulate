package app

import (
	"fmt"
	"net/http"

	"github.com/Lacka90/gophulate/processor"
	"github.com/julienschmidt/httprouter"
)

// Handle - file upload handler
func Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	mode, err := GetProcessMode(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	proc, err := GetProcessor(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintf(w, "Error: %s\n", err)
		return
	}
	content := ReadFileContent(r)

	clients := ParseCSV(content)

	p := new(processor.Processor)
	p.Comp(proc)
	message := p.Process(clients, mode)

	fmt.Fprintf(w, "%s", message)
	return
}
