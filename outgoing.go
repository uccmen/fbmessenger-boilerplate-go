package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/stvp/rollbar"
)

func handleOutgoing(w http.ResponseWriter, message Message) {
	outgoingMessage := OutgoingMessage{}
	outgoingMessage.Recipient.ID = message.Sender.ID
	outgoingMessage.Message.Text = message.MessageData.Text

	bodyB, err := json.Marshal(outgoingMessage)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", os.Getenv("FB_MESSENGER_URL"), bytes.NewBuffer(bodyB))
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	params := url.Values{}
	params.Set("access_token", os.Getenv("FB_PAGE_ACCESS_TOKEN"))

	req.URL.RawQuery = params.Encode()
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	dump, err := httputil.DumpRequestOut(req, true)
	if err != nil {
		rollbar.Error(rollbar.ERR, err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	fmt.Printf("%q", dump)

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response failed to send successfully")
		rollbar.Error(rollbar.ERR, err)
		dump, err := httputil.DumpResponse(resp, true)
		if err != nil {
			rollbar.Error(rollbar.ERR, err)
			http.Error(w, "", http.StatusInternalServerError)
			return
		}

		fmt.Printf("%q", dump)
		return
	}
}
