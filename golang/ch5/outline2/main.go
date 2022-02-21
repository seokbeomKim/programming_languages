package main

import (
	"fmt"
	"strings"
)

func square(n int) int     { return n * n }
func negative(n int) int   { return -n }
func product(m, n int) int { return m * n }

func add1(r rune) rune { return r + 1 }

func main() {
	var f func(int) int
	var g func(int, int) int

	f = square
	fmt.Println(f(3))

	f = negative
	fmt.Println(f(3))

	g = product
	fmt.Println(g(3, 4))

	fmt.Println(strings.Map(add1, "HAL-9000"))
}
