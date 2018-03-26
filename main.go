package main

import (
	"log"
	"net/http"

	"github.com/Lacka90/gophulate/app"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()
	router.POST("/process", app.Process)

	log.Fatal(http.ListenAndServe(":8080", router))
}
