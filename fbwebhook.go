package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/stvp/rollbar"
)

func fbWebhook(w http.ResponseWriter, r *http.Request) {

	dump, err := httputil.DumpRequestOut(r, true)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fmt.Printf("incoming request from webhook!! %q", dump)

	if r.Method == "GET" {
		confirmSubscription(w, r)
		return
	}

	if r.Method != "POST" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}

	// callbacks
	handleIncoming(w, r)
	return
}
