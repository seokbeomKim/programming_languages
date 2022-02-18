package main

import (
	"fmt"
	"log"
)

func deferTest() {
	fmt.Println("defer Test!")
}
func deferTest2() {
	fmt.Println("defer Test2!")
}

func main() {
	//conn, err := net.Dial("tcp", "localhost:8080")
	if true {
		log.Fatal()
	}
	//defer conn.Close()
	defer deferTest()
	defer deferTest2()
	fmt.Println("end of main function")
}
