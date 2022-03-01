package main

import (
	"io"
	"log"
	"net"
	"time"
)

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
		go handleConn(conn) // connection 처리 받는 부분을 고루틴으로 처리한다.s
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
