// Fetch prints the content found at each specified URL.
package fetcher

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

// !+
func FetchConcurrent(url string, ch chan<- string) { // Fetch prints the content found at a URL.
	start := time.Now()        // start a timer to measure the time it takes to fetch the URL
	resp, err := http.Get(url) // make an HTTP get request to the given URL and return a response and error
	if err != nil {            // if there is an error, print it and return immediately
		ch <- fmt.Sprint(err) // send to channel ch the error message
		return                // return immediately
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body) // copy the body of the response to the ioutil.Discard writer (discard everything) and return the number of bytes copied and an error
	resp.Body.Close()                                 //  close the body of the response (don't leak resources)
	if err != nil {                                   // if there is an error, print it and return immediately
		ch <- fmt.Sprintf("while reading %s: %v", url, err) // send to channel ch the error message with the URL and the error
		return                                              // return immediately
	}
	secs := time.Since(start).Seconds()                    // get the time since the start of the timer and convert to seconds
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url) // send to channel ch the time it took to fetch the URL and the number of bytes in the response and the URL
}

//!-
