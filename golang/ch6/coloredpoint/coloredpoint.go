package main

import (
	"fmt"
	"image/color"

	"geometry"
)

type MultiInheritance struct {
	v1, v2 int
}

func (p *MultiInheritance) getFirst() int {
	return p.v1
}

func (p *MultiInheritance) getSecond() int {
	return p.v2
}

// ambiguous 메서드 테스트
// func (p *MultiInheritance) ScaleBy(factor float64) float64 {
// 	return float64(p.v1) * factor
// }

type ColoredPoint struct {
	// 세 개의 필드를 갖는 대신, Point struct를 내장해서 X, Y를 사용할 수
	// 있게 했다. 이는 리눅스 커널에서 C 스타일의 상속 방식과 비슷하다.
	geometry.Point
	// mi    MultiInheritance
	MultiInheritance
	Color color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Point.Y)

	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}

	// Linux Kernel의 Embedded Structure 와 비슷한 구조를 갖고 있지만 내부
	// 구조체가 갖고 있는 메서드는 상속한 클래스가 갖고 있는 것처럼 사용할
	// 수 있다.

	// 그렇다면, 상속하는 구조체가 여러개인 경우에도 마찬가지인가? (다중 상속 가능?)
	var p = ColoredPoint{geometry.Point{1, 1}, MultiInheritance{1, 2}, red}
	var q = ColoredPoint{geometry.Point{5, 4}, MultiInheritance{3, 4}, blue}
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))

	// 가능하다! 그렇다면, 명시적으로 필드명을 정의할 경우에는 어떻게 될까?
	// 이 경우에는 inheritance 의 개념은 사라진다. 기존 C 에서 embedded
	// structure를 사용하면 상속의 개념과 데이터 관계가 명시적으로 보이지
	// 않아 해석에 어려운 부분이 있었는데 golang에서는 이러한 부분을 구분할
	// 수 있어서 좋은 것 같다.

	// fmt.Println(p.mi.getFirst())
	fmt.Println(p.getFirst())
	fmt.Println(q.getSecond())

	// 만약 다중상속하는 두 개의 클래스에서 동일 메서드명이 제공된다면 어떻게 되는가?
	// ambiguous call로 에러가 반환되며 컴파일되지 않는다
}
