package main

import (
	"net/http"
	"os"

	notify "github.com/alphagov/notifications-go-client"
	"github.com/gorilla/mux"
	"github.com/onsdigital/log.go/log"
)

var client notify.Client

func main() {
	log.Event(nil, "Starting rm-notify-gateway")

	config := notify.Configuration{
		APIKey:    []byte(os.Getenv("NOTIFY_API_KEY")),
		ServiceID: os.Getenv("NOTIFY_SERVICE_ID"),
	}

	client, err := notify.New(config)
	if err != nil {
		panic(1)
	}

	router := mux.NewRouter()
	router.HandleFunc("/info", getInfo).Methods("GET")
	router.HandleFunc("/texts/{censusUacSmsTemplateId}", sendTextMessage).Methods("POST")
	http.ListenAndServe(":8000", router)
}
