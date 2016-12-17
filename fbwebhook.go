package main

import (
	"io/ioutil"
	"log"
	"net/http"

	"github.com/bugsnag/bugsnag-go"
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

	// app.post('/webhook', function (req, res) {
	//   var data = req.body;

	//   // Make sure this is a page subscription
	//   if (data.object === 'page') {

	//     // Iterate over each entry - there may be multiple if batched
	//     data.entry.forEach(function(entry) {
	//       var pageID = entry.id;
	//       var timeOfEvent = entry.time;

	//       // Iterate over each messaging event
	//       entry.messaging.forEach(function(event) {
	//         if (event.message) {
	//           receivedMessage(event);
	//         } else {
	//           console.log("Webhook received unknown event: ", event);
	//         }
	//       });
	//     });

	//     // Assume all went well.
	//     //
	//     // You must send back a 200, within 20 seconds, to let us know
	//     // you've successfully received the callback. Otherwise, the request
	//     // will time out and we will keep trying to resend.
	//     res.sendStatus(200);
	//   }
	// });

	// function receivedMessage(event) {
	//   // Putting a stub for now, we'll expand it in the following steps
	//   console.log("Message data: ", event.message);
	// }

	reqB, err := ioutil.ReadAll(r.Body)
	if err != nil {
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}

	log.Println(string(reqB))

}
