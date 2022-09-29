package main

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
)

type LineCounter struct {
	c int
}

func (l *LineCounter) Write(p []byte) (n int, err error) {
	scanner := bufio.NewScanner(bytes.NewReader(p))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		l.c++
	}
	return len(p), nil
}

func main() {

	s := "Hello, World!\nHello, 世界！"
	var lc LineCounter

	fmt.Fprintf(&lc, s)
	fmt.Println(lc)
	fmt.Printf("no of lines %v\n", lc.c)
	net.Dial("tcp", "localhost:8080")

	fmt.Println(partialPlus(6)(3))

}

func plus(x int, y int) int {
	return x + y
}

func partialPlus(x int) func(int) int {
	return func(y int) int {
		return plus(x, y)
	}
}
