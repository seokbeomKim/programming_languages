package main

import "fmt"

func main() {
	o := 6666
	x := int64(0xdeadbeef)

	fmt.Printf("%d %[1]o %#[1]o %#[1]x\n", o)
	fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
