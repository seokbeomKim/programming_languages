package main

import (
	"bytes"
	"fmt"
	"io"
)

const debug = false

func main() {
	// var buf *bytes.Buffer
	var buf io.Writer
	if debug {
		buf = new(bytes.Buffer)
	}

	f(buf)
	if debug {
		fmt.Println("debug is on")
	}

	// io.Writer는 인터페이스
	// var w io.Writer
	// w = os.Stdout
	// w = new(bytes.Buffer)
	// w = nil
}

func f(out io.Writer) {
	if out != nil {
		out.Write([]byte("done\n"))
	}
}
