package github

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func SearchIssues(terms []string) (*IssuesSearchResult, error) {
	q := url.QueryEscape(strings.Join(terms, " "))
	resp, err := http.Get(IssuesURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	// body, err := ioutil.ReadAll(resp.Body)

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// var result []Issue
	// err = json.Unmarshal(body, &result)
	var result IssuesSearchResult

	if err := json.NewDecoder(resp.Body).Decode(&result.Items); err != nil {
		return nil, err
	}
	return &result, err
}
