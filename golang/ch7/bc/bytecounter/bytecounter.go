package bytecounter

import (
	"bufio"
	"bytes"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	// int를 Bytecounter로 변환 단순 캐스팅인데 아래와 같이 essential 타입에
	// 대해서도 객체 생성하듯이 해야하는가? 일반 캐스팅 연산은 안되는건가?

	// *c += ByteCounter(len(p))

	// 아래와 같이 단순 c 스타일의 캐스팅을 해도 동작한다. 그렇다면 위의
	// constructor 형태로 할 경우 어떤 점이 장점인가?
	*c += (ByteCounter)(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		*c++
	}
	return int(*c), nil
}
