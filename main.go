package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/onsdigital/log.go/log"
)

func main() {
	log.Event(nil, "Starting rm-notify-gateway")
	router := mux.NewRouter()
	router.HandleFunc("/info", getInfo).Methods("GET")
	router.HandleFunc("/texts/{censusUacSmsTemplateId}", sendTextMessage).Methods("POST")
	http.ListenAndServe(":8000", router)
}
