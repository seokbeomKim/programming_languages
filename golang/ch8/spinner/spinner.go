package main

import (
	"fmt"
	"time"
)

// golang에서 말하는 go routine은 Go 런타임 상에서의 스레드를 의미한다. 때문에
// 커널에서 스케쥴링되는 실제 OS 스레드와 반드시 구별해야 하며 논리적으로
// 런타임에서 구현되는 스레드를 go routine이라 한다.
// 즉, 런타임 내에서 논리적으로 구현된 OS 상위 레벨의 스레드라 할 수 있다.

func main() {
	go spinner(100 * time.Millisecond)
	const n = 45
	fibN := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
	// 여기서는 메인 함수가 종료됨에 따라 thread도 함께 종료되는 것으로
	// 나오지만, 각 go루틴이 명시적으로 스스로 종료되게 할 수도 있다.
}

func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
