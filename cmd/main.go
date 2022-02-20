package main

import (
	"log"
	"net/http"
	"time"

	"github.com/ejuju/ws-autocomplete-server/internal/ws"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(false) // init multiplexer

	router.HandleFunc("/", ws.Serve) // ws endpoint

	server := http.Server{
		Addr:              ":8080",
		Handler:           router,
		IdleTimeout:       5 * time.Second,
		ReadTimeout:       5 * time.Second,
		WriteTimeout:      5 * time.Second,
		ReadHeaderTimeout: 2 * time.Second,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
