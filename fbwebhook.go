package main

import (
	"log"
	"net/http"
)

func fbWebhook(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		confirmSubscription(w, r)
		return
	}

	// callbacks
	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	log.Println(r.Form)

}
