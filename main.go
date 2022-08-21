package main

import (
	"os"

	"mobiledatabooks.com/fetcher"
)

// go run main.go https://golang.org http://gopl.io https://godoc.org

func main() {
	for _, url := range os.Args[1:] {
		fetcher.Fetch(url)
	}
}

//!-
