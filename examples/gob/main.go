package main

import (
	"bytes"
	"encoding/gob"
)

type XPTO struct {
	Number    int
	Boolean   *bool
	Interface interface{}
	Text      *string
}

func main() {

	n := 0
	b := true
	i := 0
	t := ""

	xIn := XPTO{
		Number:    n,
		Boolean:   &b,
		Interface: &i,
		Text:      &t,
	}

	var bufIn bytes.Buffer
	enc := gob.NewEncoder(&bufIn)
	if err := enc.Encode(xIn); err != nil {
		panic(err)
	}

	var xOut XPTO

	var bufOut bytes.Buffer
	dec := gob.NewDecoder(&bufOut)
	bufOut.Write(bufIn.Bytes())
	if err := dec.Decode(&xOut); err != nil {
		panic(err)
	}

	if xOut.Text == nil {
		panic("nao deveria ser nil")
	}

}
