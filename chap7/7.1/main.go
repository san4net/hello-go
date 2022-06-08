package main

import (
	"fmt"
	"io"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

type NewWriter struct {
	w io.Writer
	c int64
}

func (n1 *NewWriter) Write(p []byte) (n int, err error) {
	n1.w.Write(p)
	n1.c += int64(len(p))
	return len(p), nil
}

func countingWriter(w io.Writer) (io.Writer, *int64) {
	nw := NewWriter{
		w: nil,
		c: 0,
	}
	return &nw, &nw.c
}

func main() {
	var sample ByteCounter
	r, err := sample.Write([]byte("hello"))
	if err != nil {

	}
	fmt.Println(r)
}
