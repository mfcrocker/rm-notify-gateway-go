package main

import (
	"encoding/json"
	"net/http"

	notify "github.com/alphagov/notifications-go-client"

	"github.com/gorilla/mux"
	"github.com/onsdigital/log.go/log"
)

type textrequest struct {
	PhoneNumber     string            `json:"phoneNumber"`
	Personalisation map[string]string `json:"personalisation"`
	Reference       string            `json:"reference"`
}

func sendTextMessage(w http.ResponseWriter, r *http.Request) {
	id, ok := mux.Vars(r)["censusUacSmsTemplateId"]
	if !ok {
		log.Event(nil, "censusUacSmsTemplateId not provided for /texts")
		http.Error(w, "Please provide a GOV.UK Notify text message template ID", http.StatusBadRequest)
		return
	}

	if r.Header.Get("Content-type") != "application/json" {
		log.Event(nil, "/texts called with invalid content type", log.Data{"id": id, "content-type": r.Header["Content-type"]})
		http.Error(w, "Invalid content type, requires application/json", http.StatusUnsupportedMediaType)
		return
	}

	var request textrequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		log.Event(nil, "Malformed JSON in body of request to /texts", log.Data{"err": log.Error(err), "id": id})
		http.Error(w, "Malformed JSON in body of request", http.StatusBadRequest)
		return
	}

	response, err := client.SendSms(request.PhoneNumber, id, request.Personalisation, request.Reference)
	if err != nil {
		if apierr, ok := err.(*notify.APIError); ok {
			log.Event(nil, "Error received from GovNotify API", log.Data{"status_code": apierr.StatusCode, "err": apierr.Message})
			http.Error(w, apierr.Message, apierr.StatusCode)
			return
		}

		log.Event(nil, "Unhandled error in GovNotify client", log.Data{"err": log.Error(err), "id": id})
		http.Error(w, "Unhandled error in GovNotify client", http.StatusInternalServerError)
		return
	}

	log.Event(nil, "SMS sent to GovNotify", log.Data{"ID": response.ID})
}
