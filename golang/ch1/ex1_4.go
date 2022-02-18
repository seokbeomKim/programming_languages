package main

import (
	"bufio"
	"fmt"
	"os"
)

// 파일을 읽어 중복되는 라인을 구해내는 예제 코드
func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			// os.Open: file descriptor, error 반환
			f, err := os.Open(arg)
			if err != nil {
				// %v: 원래 형태의 값
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)

			duplicated := false
			for _, n := range counts {
				if n > 1 {
					duplicated = true
					break
				}
			}
			f.Close()

			if duplicated == true {
				fmt.Printf("%s\n", arg)
			}
		}
	}

}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}
