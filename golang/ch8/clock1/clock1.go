package main

import (
	"io"
	"log"
	"net"
	"time"
)

// 기존 예제에서 사용했던 http 서버와 달리 tcp 서버를 예제로 사용한다.

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	// 테스트
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		handleConn(conn) // 한 번에 하나의 접속 처리
	}
}

func handleConn(c net.Conn) {
	// defer 이용해 connection 처리 이 예제의 경우는 한 연결에 대해서
	// 동기화하여 처리하므로, 동시에 여러 개의 connection 이 발생할 경우에는
	// 처리할 수 없다.
	defer c.Close()
	for {
		_, err := io.WriteString(c, time.Now().Format("15:04:05\n"))
		if err != nil {
			return
		}
		time.Sleep(1 * time.Second)
	}
}
