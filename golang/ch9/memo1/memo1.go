package main

import (
	"fmt"

	_ "../test"
)

type Memo struct {
	f     Func
	cache map[string]result
}

// Func는 기억할 함수의 타입
type Func func(key string) (interface{}, error)
type result struct {
	// interface{} means you can put value of any type, including your own
	// custom type.
	value interface{}
	err   error
}

func New(f Func) *Memo {
	return &Memo{f: f, cache: make(map[string]result)}
}

func (memo *Memo) Get(key string) (interface{}, error) {
	res, ok := memo.cache[key]
	if !ok {
		res.value, res.err = memo.f(key)
		memo.cache[key] = res
	}
	return res.value, res.err
}

func main() {
	fmt.Println("memo1")
}
