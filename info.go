package main

import (
	"encoding/json"
	"net/http"

	"github.com/onsdigital/log.go/log"
)

type appinfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func info() appinfo {
	return appinfo{Name: "rm-notify-gateway", Version: "0.0.1"}
}

func getInfo(w http.ResponseWriter, r *http.Request) {
	log.Event(nil, "Responding to /info", log.Data{"ip": r.RemoteAddr})
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(info())
	if err != nil {
		log.Event(nil, "Couldn't return info", log.Error(err))
		http.Error(w, "Couldn't return info", http.StatusInternalServerError)
	}
}
