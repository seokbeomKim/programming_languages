package main

import (
	"fmt"

	"./bytecounter"
)

func main() {
	// [NOTICE] 패키지 내에 정의되어 있는 것 사용시 반드시
	// {패키지명}.{something}의 형태로 사용해야 한다.
	var c bytecounter.ByteCounter
	// var lc bytecounter.LineCounter
	// var wc bytecounter.WordCounter

	c.Write([]byte("Hello"))
	fmt.Println(c)

	c = 0
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)

	fmt.Fprintf(&c, "hello, %s. The word count would be..", name)
	fmt.Println(c)
}
