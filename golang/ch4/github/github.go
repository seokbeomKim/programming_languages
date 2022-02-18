package github

import "time"

const IssuesURL = "https://api.github.com/repos/seokbeomKim/TistoryWriter/issues"

// const IssuesURL = "https://reqres.in/api/users/2"

type IssuesSearchResult struct {
	Items []*Issue
}

type Issue struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}
