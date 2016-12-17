package main

import (
	"log"
	"net/http"
	"os"

	"github.com/bugsnag/bugsnag-go"
)

func main() {
	http.HandleFunc("/health", healthCheck)

	err := http.ListenAndServe(":"+os.Getenv("PORT"), bugsnag.Handler(nil))
	if err != nil {
		log.Panicln("ListenAndServe: ", err)
	}
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "", http.StatusMethodNotAllowed)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("I'm OK!"))
	if err != nil {
		bugsnag.Notify(err)
		http.Error(w, "", http.StatusInternalServerError)
		return
	}
}
