package main

import (
	"bufio"
	"ch2/tempconv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args[1:]) == 0 {
		input := bufio.NewScanner(os.Stdin)
		for input.Scan() {
			tryConv(input.Text())
		}
	} else {
		for _, arg := range os.Args[1:] {
			tryConv(arg)
		}
	}
}

func tryConv(v string) {
	t, err := strconv.ParseFloat(v, 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cf: %v\n", err)
		os.Exit(1)
	}
	f := tempconv.Fahrenheit(t)
	c := tempconv.Celsius(t)
	fmt.Printf("%v = %v, %v = %v\n", f, tempconv.FToC(f), c, tempconv.CToF(c))
}
