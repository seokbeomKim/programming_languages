package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	// 배열 변수 초기화 방법 [...]를 이용하나 []을 이용하나 초기화 값의
	// 개수가 결정되는데 다른 점이 있는가?
	var q [3]int = [3]int{1, 2, 3}
	var p []int = []int{1, 2, 3, 4}
	r := [...]int{1, 2, 3, 4}
	t := []int{1, 2, 3, 4}

	fmt.Println(q)
	fmt.Println(p)
	fmt.Println(r)
	fmt.Println(t)

	// 아래 타입을 비교하면 그 차이점을 알 수 있다.
	// [4]int vs. []int
	// 고정 크기의 배열 vs. 슬라이스
	fmt.Printf("%T vs. %T\n", r, t)

	// 앞서 정의한 const 값을 출력해보자
	fmt.Println(USD)
	fmt.Println(EUR)
	fmt.Println(GBP)
	fmt.Println(RMB)

	symbol := [...]string{USD: "달러", EUR: "유로"}
	fmt.Println(symbol)
}
