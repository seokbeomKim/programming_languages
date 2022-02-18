package xkcd

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const (
	xkcd_url = "https://xkcd.com/%d/info.0.json"
)

type XkcdInfo struct {
	Num        int
	Title      string
	SafeTitle  string `json: "safe_title"`
	Img        string
	Transcript string
}

func SearchKeyword(keyword *string) ([]*XkcdInfo, error) {
	// 단순하게 json 파일을 탐색할 수도 있지만 전체 JSON 파일을 특정 데이터
	// 타입으로 unmarshaling 한 후에 키워드를 검색하는 방식으로 구현하도록
	// 한다.
	files, err := os.ReadDir("./downloaded")
	if err != nil {
		return nil, err
	}

	var infoSet []*XkcdInfo

	for _, file := range files {
		var parsed = new(XkcdInfo)
		path := fmt.Sprintf("./downloaded/%s", file.Name())
		data, err := ioutil.ReadFile(path)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read file %s", file.Name())
			return nil, err
		}
		json.Unmarshal(data, parsed)
		infoSet = append(infoSet, parsed)
	}

	var rvalue []*XkcdInfo

	for _, item := range infoSet {
		if strings.Contains(item.Title, *keyword) {
			rvalue = append(rvalue, item)
		}
	}

	return rvalue, nil
}

func Download(index *int) {
	var filename, url string

	for i := 1; i < *index; i++ {
		filename = fmt.Sprintf("./downloaded/%d.json", i)
		url = fmt.Sprintf(xkcd_url, i)
		downloadFile(filename, url)
	}
}

func downloadFile(filename string, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	out, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}
