package main

import (
	"fmt"
)

func main() {
	// 고루틴 채널 생성을 위해서는 반드시 make + chan (type) 키워드를 이용한다.
	naturals := make(chan int)
	squares := make(chan int)

	// 현재 예제는 Counter -> Squarer -> Printer 로 이루어지는 3개의
	// 고루틴을 채널을 이용하여 단방향으로 연결하는 방법을 예로 든다. 어느
	// 하나의 counter - squarer - printer 의 for 문이 없어져 스레드 하나라도
	// 미리 끝난다면, deadlock 이 발생 가능하다.

	// Counter
	go func() {
		for x := 0; ; x++ {
			naturals <- x
		}
	}()

	// Squarer
	go func() {
		for {
			x := <-naturals
			squares <- x * x
		}

	}()

	// Printer
	for {
		fmt.Println(<-squares)
		// time.Sleep(time.Second)
	}
}
