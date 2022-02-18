package main

import "fmt"

func main() {
	s := [...]string{"d", "d", "d", "run", "please", "run", "run"}
	r := removeDuplicate(s[:])
	fmt.Printf("%v\n", s)
	fmt.Printf("%v\n", r)
}

func removeDuplicate(arr []string) []string {
	last := ""
	offset := 0

	for i, s := range arr {
		if s == "" {
			continue
		}
		if last == s {
			offset++
		}
		fmt.Println(i - offset)

		last = s
		arr[i-offset] = s
	}
	return arr[:len(arr)-offset]
}
