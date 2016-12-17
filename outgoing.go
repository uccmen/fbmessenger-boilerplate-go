package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"

	"github.com/bugsnag/bugsnag-go"
)

func handleOutgoing(w http.ResponseWriter, message Message) {
	outgoingMessage := OutgoingMessage{}
	outgoingMessage.Recipient.ID = message.Sender.ID
	outgoingMessage.Message.Text = message.MessageData.Text

	bodyB, err := json.Marshal(outgoingMessage)
	if err != nil {
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest("POST", os.Getenv("FB_MESSENGER_URL"), bytes.NewBuffer(bodyB))
	if err != nil {
		bugsnag.Notify(err)
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
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		err = fmt.Errorf("Response failed to send successfully")
		bugsnag.Notify(err, resp)
		log.Println(err.Error())
		return
	}
}
