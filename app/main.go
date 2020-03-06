package main

import (
	"fmt"
	"handlers/handlers"
	"log"
	"net/http"
	"os"
	"slack/config"
	"strings"
)

func init() {
	s := config.GetSlackItem()
	http.HandleFunc("/v1/event-point", func(w http.ResponseWriter, r *http.Request) {
		handlers.EventPoint(w, r, s)
	})
}

func main() {
	port := []string{":", os.Getenv("APP_PORT")}
	fmt.Println("[INFO] Server listening")
	if err := http.ListenAndServe(strings.Join(port, ""), nil); err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
}
