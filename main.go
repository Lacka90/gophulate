package main

import (
	"log"
	"net/http"

	"github.com/Lacka90/gophulate/app"
	"github.com/julienschmidt/httprouter"
)

func main() {
	new(App).Init().Run(":8080")
}

// App - struct
type App struct {
	router *httprouter.Router
}

// Init - func
func (a *App) Init() *App {
	a.router = httprouter.New()
	a.router.POST("/process", app.Handle)
	return a
}

// Run - func
func (a *App) Run(address string) {
	log.Fatal(http.ListenAndServe(address, a.router))
}
