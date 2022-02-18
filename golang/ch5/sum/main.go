package main

import (
	"errors"
	"fmt"
)

func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func min(vals ...int) (int, error) {
	if len(vals) == 0 {
		return 0, errors.New("At least two numbers required")
	}
	min := vals[0]
	for _, val := range vals[1:] {
		if val < min {
			min = val
		}
	}
	return min, nil
}

func f(...int) {}
func g([]int)  {}

func main() {
	// 아래와 같이 사용할 때 함수의 vals 타입은 []int 슬라이스이다.
	// 하지만 이러한 가변 인자 함수의 타입은 일반 슬라이스 파라미터를 받는 함수의 타입과 다르다.
	fmt.Println(sum())
	fmt.Println(sum(3))
	fmt.Println(sum(1, 2, 3, 4))

	fmt.Printf("%T\n%T\n", f, g)

	fmt.Println(min())
}
