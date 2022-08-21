// Fetch prints the content found at each specified URL.
package fetcher

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

//!+

func Fetch(url string) { // Fetch prints the content found at a URL.
	resp, err := http.Get(url) // make an HTTP get request to the given URL
	if err != nil {            // if there is an error, print it and return  immediately
		fmt.Fprintf(os.Stderr, "fetch: %v\n", err) // print error to stderr
		os.Exit(1)                                 // exit with error code 1 (error)
	}
	b, err := io.Copy(os.Stdout, resp.Body) // copy the body of the response to the stdout and return the number of bytes copied and an error (if there is one)
	resp.Body.Close()                       // close the body of the response (don't leak resources)
	if err != nil {                         // if there is an error, print it and return immediately
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err) // print error to stderr
		os.Exit(1)                                                  // exit with error code 1 (error)
	}
	fmt.Printf("%s", b) // print the body of the response to stdout (the terminal)
}

//!-
