// Fetch prints the content found at each specified URL.
package fetcher

// MIT License

// Copyright (c) 2022 Mobile Data Books, LLC

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

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

	fmt.Println("HTTP status:", resp.Status)

	_, err = io.Copy(io.Discard, resp.Body) // copy the body of the response to the stdout and return the number of bytes copied and an error (if there is one)
	resp.Body.Close()                       // close the body of the response (don't leak resources)
	if err != nil {                         // if there is an error, print it and return immediately
		fmt.Fprintf(os.Stderr, "fetch: reading %s: %v\n", url, err) // print error to stderr
		os.Exit(1)                                                  // exit with error code 1 (error)
	}
	// fmt.Printf("%s", b) // print the body of the response to stdout (the terminal)
}

//!-
