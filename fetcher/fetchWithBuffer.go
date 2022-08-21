// Fetch prints the content found at each specified URL.
package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//!+

func FetchWithBuffer(url string) { // Fetch prints the content found at a URL.
	resp, err := http.Get(url) // make an HTTP get request to the given URL
	if err != nil {            // if there is an error, print it and return  immediately
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // print error to stderr
		os.Exit(1)                                 // exit with error code 1 (error)
	}
	fmt.Println("HTTP status:", resp.Status)

	_, err = io.ReadAll(resp.Body) // read the body of the response and return the bytes and an error (if there is one)
	resp.Body.Close()              // close the body of the response (don't leak resources)
	if err != nil {                // if there is an error, print it and return immediately
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err) // print error to stderr
		os.Exit(1)                                                  // exit with error code 1 (error)
	}
	// fmt.Printf("%s", b) // print the body of the response to stdout (the terminal)
}

//!-
