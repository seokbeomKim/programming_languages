package bytecounter

import (
	"bufio"
	"bytes"
)

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		*c++
	}
	return int(*c), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	buf := bytes.NewBuffer(p)
	scanner := bufio.NewScanner(buf)
	for scanner.Scan() {
		*c++
	}
	return int(*c), nil
}
