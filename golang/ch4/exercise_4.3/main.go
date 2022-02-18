package main

import "fmt"

func main() {
	// [...]의 의미는 무엇인가?
	// ...는 varadic 을 의미한다. 즉, 배열의 크기를 전달한 원소에 따라
	// 결정되도록 선언한다.
	a := [...]int{0, 1, 2, 3, 4, 5, 6}
	reverse(&a)
	fmt.Println(a)
}

func reverse(a *[7]int) {
	for i, j := 0, len(*a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}
