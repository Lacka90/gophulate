package app

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Handle - file upload handler
func Handle(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "OK")
}
