package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func test_func(data []byte) {
	fmt.Printf("%T\n", data)
}

func main() {

	var r []byte

	// golang에서는 function pointer를 지원하지 않는다.
	// https://golangdocs.com/pointer-to-a-function-in-golang
	switch os.Args[1] {
	case "sha256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(os.Args[2])))
		break
	case "sha384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(os.Args[2])))
		break
	case "sha512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(os.Args[2])))
		break
	}

	fmt.Printf("return hash: %v\n", r)

	c1 := sha256.Sum256([]byte("y"))
	c2 := sha256.Sum256([]byte("x"))

	fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)

	// zero(&c1)
	cnt := compare(&c1, &c2)
	fmt.Printf("compare: diff count is %d\n", cnt)
}

func zero(ptr *[32]byte) {
	// type 1
	for i := range ptr {
		ptr[i] = 0
	}

	// type 2
	*ptr = [32]byte{}
}

func compare(c1, c2 *[32]byte) int {
	fmt.Printf("c1: %x, c2: %x\n", c1, c2)

	r := 0
	for i, v := range c1 {
		if v != c2[i] {
			r += 1
		}
	}
	return r
}
