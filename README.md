## Why use lot thread when one do trick?

This project is a semi-practical experiment aiming to create an asynchronous, single-threaded HTTP server using primitive coroutines in Go, despite the pre-existing Goroutines already existing. 

Pushing the boundaries of Go's concurrency model inwards, am I right?

### Usage

1. Initiate the HTTP server by running `go run main.go`, the server will be listening on port 8080.
2. Submit requests to the server, a number after the slash means the request will take that many seconds to complete due to the server's artificial CONCURRENT delay.

```bash
for i in {10..0}; do
  curl "localhost:8080/${i}" &
done
wait
```

