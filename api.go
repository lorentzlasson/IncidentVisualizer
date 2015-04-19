package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

func handleApiDefault(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "APIs:")
	fmt.Fprintf(w, "/hw")
	fmt.Fprintf(w, "/bc")
}

func handleApiHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func handleApiBroadcast(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "broadcast")
	queryParams := r.URL.Query()
	message := queryParams.Get("msg")
	h.broadcast <- []byte(message)
}

func setupRoutes() {
	router := mux.NewRouter()
	subRouter := router.PathPrefix("/api").Subrouter().StrictSlash(false)
	subRouter.HandleFunc("/", handleApiDefault)
	subRouter.HandleFunc("/hw", handleApiHelloWorld)
	subRouter.HandleFunc("/bc", handleApiBroadcast).Methods("POST")

	router.PathPrefix("/ws").HandlerFunc(serveWs)
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("./static/")))

	http.Handle("/", router)
}
