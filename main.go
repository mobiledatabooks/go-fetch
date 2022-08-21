package main

import (
	"os"

	"github.com/mobiledatabooks/go-fetch/fetcher"
)

// https://github.com/mobiledatabooks/go-fetch.git

// go run main.go https://golang.org http://gopl.io https://godoc.org

func main() {
	for _, url := range os.Args[1:] {
		fetcher.Fetch(url)
	}
}

//!-
