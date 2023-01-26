package main

import (
	"fmt"
)

type ExampleStruct struct {
	Text string
}

// New title
// @A provide instance=github.com/americanas-go/inject/examples/simple.ExampleStruct name=xpto
func New() *ExampleStruct {
	return &ExampleStruct{
		Text: "Hello World",
	}
}

// Xpto title
// @A inject instance=github.com/americanas-go/inject/examples/simple.ExampleStruct name=xpto
// @A invoke
func Xpto(ex *ExampleStruct) {
	fmt.Printf("invoked: %s", ex.Text)
}
