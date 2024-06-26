# Concurrent Ping-Pong Server and Client

## Overview
I was having fun implementing a concurrent Ping-Pong server and client using Go that communicate over simple TCP sockets. The server can handle multiple client connections simultaneously, broadcasting "pong" messages to all connected clients at random intervals. The clients connect to the server, send "ping" messages at random intervals, and print out "pong" messages received from the server.


## Benchmark Test
There is also a benchmark that I did to check whether casting a byte array to a string for comparison or using `bytes.Equal` was faster. It turns out that for my size of messages, they perform the same.

```
BenchmarkBytesEqual-12                  1000000000               0.1280 ns/op
BenchmarkBytesToStringEqual-12          1000000000               0.1176 ns/op
```

## Run It

* Have Go installed
* Open a terminal tab or window and run `make pong`.
* Open another terminal tab or window and run `make ping` (This will spawn 1000 clients).
* `make bench`if you want to try the benchmarks

## Conclusion
It was a nice project to play around with transmitting binary data over TCP, dealing with threads (goroutines, actually), locks, wait groups, channels, and all the cool concurrency Go stuff.