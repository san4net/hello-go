package main

import (
	"flag"
	"fmt"
	"hello-go/chap7/7.6/tempconv"
)

var temp = tempconv.CelsiusFlag("temp", 20.0, "the temperature")

func main() {
	flag.Parse()
	fmt.Print(*temp)

	var k Keyboard = computer{}
	k.keyboardType()
}

type Keyboard interface {
	keyboardType() string
	CommandToPrint(value string) bool
}

type computer struct {
	name     string
	Keyboard Keyboard
}

func (c computer) keyboardType() string {
	return "querty keyboard"
}

func (c computer) CommandToPrint(values string) bool {
	return true
}
