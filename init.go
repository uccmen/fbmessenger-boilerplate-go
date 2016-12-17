package main

import (
	"log"
	"os"
)

func init() {
	if os.Getenv("PORT") == "" {
		log.Panicln("PORT not set")
	}
	if os.Getenv("HUB_VERIFY_TOKEN") == "" {
		log.Panicln("HUB_VERIFY_TOKEN not set")
	}
}
