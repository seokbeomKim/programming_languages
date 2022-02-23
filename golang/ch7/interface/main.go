package main

import (
	"bytes"
	"io"
	"os"
	"time"
)

type Artifact interface {
	Title() string
	Creator() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Streamer
	Format() string
}

type Video interface {
	Streamer
	Resolution() (x, y int)
}

type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

func main() {
	// io.Writer는 Write 메서드를 구현하는 인터페이스이다.
	var w io.Writer

	// os.Stdout은 os.File 로서 io.Writer 인터페이스를 만족하기 위해서는
	// Write 메서드가 구현되어 있어야 한다.
	w = os.Stdout

	// (*Buffer) 타입 또한 Write 메서드를 구현하고 있으므로 인터페이스 변수로
	// assign 할 수 있다.
	w = new(bytes.Buffer)

	// const 는 할당 불가
	w = time.Second

	// Reader, Writer, Closer 인터페이스를 상속하여 만든 인터페이스
	var rwc io.ReadWriteCloser
	rwc = os.Stdout

	// Close가 없으므로 할당 불가
	rwc = new(bytes.Buffer)

	// Close 제외한 나머지 인터페이스는 구현되어 있으므로 아래와 같이 할당
	// 가능하다.
	var rw io.ReadWriter
	rw = new(bytes.Buffer)

	w = rwc
	rwc = w
	w = rw

	// (*bytes).Buffer는 io.Writer를 충족해야 한다.
	var _ io.Writer = (*bytes.Buffer)(nil)
}
