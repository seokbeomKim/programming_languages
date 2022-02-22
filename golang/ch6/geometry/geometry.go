package geometry

import (
	"fmt"
	"math"
)

type Point struct{ X, Y float64 }
type Path []Point
type P *int // 포인터의 경우 메서드 정의가 불가능하다.

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

// 전통 함수
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// 동일한 역할을 수행하는 Point 타입의 메서드
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64) {
	p.X *= factor
	p.Y *= factor
}

// 추가 파라미터 p를 메서드의 수신자라 부르는데, 이것은 초기 객체지향 언어에서
// 메서드 호출을 "객체에 메시지를 전송한다" 라고 하던 관례에 따른 것이다.

func main() {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))
}
