package main

import (
	"bytes"
	"testing"
)

func bytesToStringEqual(a, b []byte) bool {
	return string(a) == string(b)
}

func BenchmarkBytesEqual(b *testing.B) {
	a := []byte("ping")
	c := []byte("ping")
	for i := 0; i < b.N; i++ {
		bytes.Equal(a, c)
	}
}

func BenchmarkBytesToStringEqual(b *testing.B) {
	a := []byte("ping")
	c := []byte("ping")
	for i := 0; i < b.N; i++ {
		bytesToStringEqual(a, c)
	}
}

func main() {
	testing.Benchmark(BenchmarkBytesEqual)
	testing.Benchmark(BenchmarkBytesToStringEqual)
}
