package main

import "fmt"

// 익명 함수를 반환하는데 var x int 값은 static 처럼 유지된다.?

func squares() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}

// 만약 anonymous function 이 사용되지 않는 경우에도 closure 개념이 포함될까?
// closure 개념은 포함되지 않는다.
func testFunc() int {
	var x int
	x++
	return x
}

func main() {
	f := squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// 만약 squares 로부터 얻은 f 로 인한 레퍼런스 카운트 때문에 var x 의
	// 값이 유지되는 것이라면, 변수 f가 해제된다면 값은 초기화되는가?
	f = nil

	f = squares()
	fmt.Println(f())
	fmt.Println(f())
	fmt.Println(f())

	// 예상했던 대로 function pointer 레퍼런스 카운터 값을 0로 만드니 기존에
	// 유지되었던 variable에 대한 값이 유지되지 않았다. golang에서도 얘기한
	// closure 개념(function variable을 이용한)은 이러한 레퍼런스 값을
	// 이용한 트릭?에 가까운 느낌이다.
}
