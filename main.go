package main

import (
	"os"

	"github.com/mobiledatabooks/go-fetch/fetcher"
)

// https://github.com/mobiledatabooks/go-fetch.git

// go run main.go https://golang.org http://gopl.io https://godoc.org

func main() { // Fetch prints the content found at each specified URL.
	for _, url := range os.Args[1:] { // for each URL in the command line arguments
		fetcher.FetchWithBuffer(url) //  fetch the URL and print the content (with buffer)
	}
}

//!-
