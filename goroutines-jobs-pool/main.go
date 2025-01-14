package main

import (
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Code int
}

func main() {
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results) //Notice how I start the routines before the list is ready.
	}

	urls := []string{
		"https://google.com",
		"https://google.com",
		"https://google.com",
		"https://google.com",
		"https://youtube.com",
	}

	for _, url := range urls {
		jobs <- Site{url} //Fills the queue only if buffer not full. Otherwise, stops the main process and waits.
	}
	close(jobs) //If I don't close jobs, the crawls goroutines will remain pending on "range jobs"

	for a := 0; a < len(urls); a++ {
		result := <-results //This stops the main process until it receives something
		log.Println(result)
	}
}

func crawl(wId int, jobs chan Site, results chan Result) {
	for site := range jobs {
		log.Printf("Worker Id: %d", wId)
		resp, _ := http.Get(site.URL)

		results <- Result{Code: resp.StatusCode}
	}
}
