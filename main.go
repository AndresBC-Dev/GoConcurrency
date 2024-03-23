package main

import (
	"log"
	"net/http"
	"sync"
)

var urls = []string{
	"http://localhost:1234?duration=3s",
	"http://localhost:1234?duration=1s",
	"http://localhost:1234?duration=5s",
}

func main() {
	fetchConcurrent(urls)
}

func fetchSecuential(ursl []string) {
	for _, url := range ursl {
		fetch(url)
	}
}

func fetchConcurrent(urls []string) {
	var wg sync.WaitGroup
	wg.Add(3)

	for _, url := range urls {
		go func() {
			fetch(url)
			wg.Done()
		}()

	}
	wg.Wait()
}

func fetch(url string) {
	resp, err := http.Head(url)
	if err != nil {
		log.Fatalf("failed url: %S, err: %V", url, err)
	}
	log.Println(url, ": ", resp.StatusCode)
}
