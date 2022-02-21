package main

import "fmt"

func main() {
	// 스코프로 인해 발생하는 변수의 값 문제에 대한 예제
	total := 0

	var dir []int
	var fs []func()

	for i := 0; i < 10; i++ {
		total += 1
		f := func() {
			dir = append(dir, i+total)
		}
		f()

		fs = append(fs, func() {
			dir = append(dir, dir[-1])
		})
	}
	fmt.Println(dir)
	for _, f := range fs {
		f()
	}
	fmt.Println(dir)
}
