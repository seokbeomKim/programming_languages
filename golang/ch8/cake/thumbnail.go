package thumbnail

import "log"

// 동시성 패턴을 살펴보기 위한 예제

func ImageFile(infile string) (string, error)

func makeThumbnails(filenames []string) {
	for _, f := range filenames {
		if _, err := thumbnail.ImageFile(f); err != nil {
			log.Println(err)
		}
	}
}

func makeThumbnails2(filenames []string) {
	ch := make(chan struct{})
	for _, f := range filenames {
		go func(f string) {
			thumbnail.ImageFile(f)
			ch <- struct{}{}
		}(f)
	}
	// go-routine 이 끝날 때까지 대기
	for range filenames {
		<-ch
	}
}
