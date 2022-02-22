package main

import (
	"fmt"
	"sync"
)

// c 에서도 이러한 초기화 방법은 지원한다.
var cache = struct {
	sync.Mutex
	mapping map[string]string
}{
	mapping: make(map[string]string),
}

func main() {
	fmt.Println(Lookup("test"))
}

// 함수 안에 내부 함수 정의가 가능한가?
func Lookup(key string) string {
	cache.Lock()
	v := cache.mapping[key]
	cache.Unlock()
	return v
}
