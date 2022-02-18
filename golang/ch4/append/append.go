package main

import "fmt"

func appendInt(x []int, y int) []int {
	var z []int
	zlen := len(x) + 1
	if zlen <= cap(x) {
		// 슬라이스 확장
		z = x[:zlen]
	} else {
		// 어째서 append 하는데 make로 늘려줘야 하는가?
		// 자동으로 안되나?
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		copy(z, x)
	}
	z[len(x)] = y
	// 만약 레퍼런스 카운트가 없는 경우, 기존 배열은 메모리 해제 되는가?
	return z
}

func main() {
	// x := []int{1, 2, 3, 4}
	// x = appendInt(x, 4)
	// fmt.Printf("%d\n", x)

	var x, y []int
	for i := 0; i < 10; i++ {
		y = appendInt(x, i)
		fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
		x = y
	}
}
