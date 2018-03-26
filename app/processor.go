package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Process - file upload processor
func Process(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "OK")
}
