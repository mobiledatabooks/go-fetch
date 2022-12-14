# go-fetch

The Go programming language provides a collection of packages under the "net" umbrella that make it easy to send and receive information over the internet, establish low-level network connections, and set up servers. Go's concurrency features are particularly well-suited for this type of work.

The fetch program retrieves the content of each specified URL and prints it as uninterpreted text. This program is inspired by the invaluable utility curl.

FetchWithBuffer
io.ReadAll reads from r until an error or EOF and returns the data it read. A successful call returns err == nil, not err == EOF. Because ReadAll is defined to read from src until EOF, it does not treat an EOF from Read as an error to be reported.

Starts with 512 bytes buffer and adds more capacity if needed.

Fetch
The io.Copy function reads from a source (src) and writes to a destination (dst) until either EOF is reached on src or an error occurs. It returns the number of bytes copied and the first error encountered while copying, if any.

It can be used instead of io.ReadAll to copy the response body to os.Stdout without requiring a buffer large enough to hold the entire stream. 

FetchConcurrent
A goroutine is a function that can be executed concurrently with other functions. A channel is a communication mechanism that allows one goroutine to pass values of a specified type to another goroutine. 