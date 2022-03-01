package main

import (
	"fmt"
	"time"
)

func launch() {
	fmt.Println("Launch!!")
}

func main() {
	abort := make(chan struct{}, 1)

	fmt.Println("Commencing countdown. Press return to abort.")
	tick := time.Tick(1 * time.Second)
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)
		select {
		case <-tick:
			// nothing
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}
	launch()
}
