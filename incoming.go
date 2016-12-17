package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/bugsnag/bugsnag-go"
)

func handleIncoming(w http.ResponseWriter, r *http.Request) {
	reqB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	incomingMessage := IncomingMessage{}
	err = json.Unmarshal(reqB, &incomingMessage)
	if err != nil {
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	if incomingMessage.Object != "page" {
		http.Error(w, "only allowed to chat via fb page", http.StatusForbidden)
		return
	}

	if incomingMessage.Entries == nil {
		bugsnag.Notify(fmt.Errorf("entry is not provided"))
		http.Error(w, "", http.StatusExpectationFailed)
		return
	}

	for _, entry := range *incomingMessage.Entries {
		for _, message := range entry.Messaging {
			handleOutgoing(w, message)
		}
	}

	return
}
