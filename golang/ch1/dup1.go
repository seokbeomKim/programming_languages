package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// map 객체 생성; string 키, int 값
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
