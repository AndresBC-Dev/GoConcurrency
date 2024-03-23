package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":1234", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	duration := r.URL.Query().Get("duration")
	durationParsed, err := time.ParseDuration(duration)
	if err != nil {
		http.Error(w, "Invalid duration", http.StatusBadRequest)
		return
	}

	time.Sleep(durationParsed)
	fmt.Fprintf(w, "Request processed after %s\n", duration)
}
