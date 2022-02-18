package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"./xkcd"
)

func main() {
	// 기능 1. 0부터 입력한 번호까지 다운로드 하여 json 파일들을 특정
	// 디렉토리에 다운로드 한다.
	ptrIdx := flag.Int("idx", 10, "maximum index to download from xkcd")
	ptrKeyword := flag.String("keyword", "foo", "search keyword")

	// 기능 2. 다운로드 되어 있는 json 파일들을 검색하여 입력한 검색어로
	// 찾을 수 있는 json 파일의 URL과 대본을 출력한다.

	flag.Parse()

	fmt.Println(*ptrIdx)
	fmt.Println(*ptrKeyword)

	if _, err := os.Stat("./downloaded/"); errors.Is(err, os.ErrNotExist) {
		err = os.Mkdir("./downloaded/", 0755)
		xkcd.Download(*&ptrIdx)
	}
	items, err := xkcd.SearchKeyword(ptrKeyword)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to search keyword")
		os.Exit(1)
	}

	for _, item := range items {
		fmt.Println((*item).Img)
		fmt.Println((*item).Transcript)
	}
}
