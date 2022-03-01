package main

var (
	sema    = make(chan struct{}, 1)
	balance int
)

func Deposit(amount int) {
	// struct {}{}로 하는 이유는 무엇인가?
	// channel이 특정 데이터 없이 동기화를 위해 사용되는 것을 나타내기 위한 것
	sema <- struct{}{}
	balance = balance + amount
	<-sema // 토큰 해제
}

func Balance() int {
	sema <- struct{}{}
	b := balance
	<-sema
	return b
}
