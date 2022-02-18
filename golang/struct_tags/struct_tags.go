package main

import (
	"fmt"
)

type User struct {
	Name string `example:"name"`
}

func (u *User) String() string {
	return fmt.Sprintf("Hi! My name is %s", u.Name)
}

func main() {
	u := &User{
		Name: "Sammy",
	}
	fmt.Println(u)

	example.testStructTags(u)
}

// struct tags 를 사용하기 위해서는 또다른 go 코드가 반드시 필요하다. 해당
// 코드는 런타임에 구조체를 판단한다.
