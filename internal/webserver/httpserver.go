package webserver

import (
	"log"
	"net/http"

	repository "santcum/internal/repo"

	"github.com/gorilla/mux"
)

func StartServer() {
	log.Println("Starting the web server to handle the api requests")
	MuxRouter := mux.NewRouter()

	MuxRouter.HandleFunc("/items", repository.GetAllItems).Methods("GET")
	MuxRouter.HandleFunc("/items/{id}", repository.GetItem).Methods("GET")
	MuxRouter.HandleFunc("/create", repository.CreateItem).Methods("POST")
	http.ListenAndServe(":8080", MuxRouter)
}
