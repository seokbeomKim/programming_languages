package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Launched Nuclear Missle!")
}

func main() {
	fmt.Println("Commencing countdown.")
	// time.Tick 은 chan time.Time 타입이다!
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		// 채널 입력 대기
		<-tick
	}
	launch()
}
