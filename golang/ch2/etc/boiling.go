package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	fmt.Printf("boiling point = %gF or %gC\n", f, fToC(f))
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
