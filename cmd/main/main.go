package main

import (
	"log"
	"net/http"

	"github.com/zakhaev26/pteromys/internal/kafka"
	"github.com/zakhaev26/pteromys/internal/ws"
)

func main() {
	setupAPI()

	go func() {
		for {
			kafka.PushToKafka("deez", "saswat has dangling balls")
			log.Printf("produced message")
		}
	}()
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setupAPI() {

	manager := ws.NewManager()

	http.Handle("/", http.FileServer(http.Dir("./frontend")))
	http.HandleFunc("/ws", manager.ServeWS)

}
