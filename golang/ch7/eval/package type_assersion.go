package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	// 타입 단언 예제. 단언이라는 용어가 거슬리지만, 영어로 하면 Type
	// Assertion 이다.
	var w io.Writer
	w = os.Stdout

	// f := w.(*os.File)
	// c := w.(*bytes.Buffer)

	// x.(T) 형태로 구성하며 x는 인터페이스의 타입, T는 단언 타입이다. 즉,
	// w가 (T)로 전달되는 인터페이스를 만족하는지 여부를 확인할 수 있다.

	fmt.Println("Type assertion")
	// fmt.Println(f, c)
	fmt.Println(w.(*os.File))
	if err := w.(*os.File); err != nil {
		fmt.Println("error is not nil", err)
		log.Fatal(err)
	} else {
		fmt.Println("error is nil")
	}
}
