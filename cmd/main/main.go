package main

import (
	"log"
	"net/http"

	"github.com/zakhaev26/pteromys/internal/wsManager"
)

func main() {
	setupAPI()

	// Serve on port :8080, fudge yeah hardcoded port
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// setupAPI will start all Routes and their Handlers
func setupAPI() {
	manager := wsManager.NewManager()

	http.HandleFunc("/ws", manager.ServeWS)
}
