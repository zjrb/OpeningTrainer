package main

import (
	"backend/internal/handlers"
	"backend/internal/websocket"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/hello", handlers.Hello).Methods("GET")

	go websocket.StartServer()
	log.Println("Starting server on :8080...")

	http.ListenAndServe(":8080", r)

}
