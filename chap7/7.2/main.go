package main

import (
	"fmt"
	"io"
)

type ByteCounter int

func (b *ByteCounter) Write(p []byte) (n int, err error) {
	return len(p), nil
}

// it write and keep count of characters written
type CountingWriter struct {
	writer io.Writer
	count  int64
}

func (n1 *CountingWriter) Write(p []byte) (n int, err error) {
	n1.writer.Write(p)
	n1.count += int64(len(p))
	return len(p), nil
}

func countingWriter(w io.Writer) (io.Writer, *int64) {
	nw := CountingWriter{
		writer: w,
		count:  0,
	}
	return &nw, &nw.count
}

func main() {
	var sample ByteCounter
	r, err := sample.Write([]byte("hello"))

	if err != nil {

	}
	fmt.Println(r)

	w, count := countingWriter(&sample)
	fmt.Fprint(w, "hello again")
	fmt.Fprint(w, "hello again")
	fmt.Println(*count)
	*count = 0
	fmt.Fprint(w, "hello again")
	fmt.Println(*count)
}
