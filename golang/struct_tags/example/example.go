package main

import "fmt"

type T struct {
	f1     string "f one"
	f2     string
	f3     string `f three`
	f4, f5 int64  `f four and five`
}

func main() {
	a := T{f1: "as", f2: "qw"}
	fmt.Println(a)
}
