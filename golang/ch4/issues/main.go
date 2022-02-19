package main

import (
	"issueshtml"
	"log"
	"os"
	"text/template"

	"github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(time.Now())
	// for _, item := range result.Items {
	// 	fmt.Printf("#%-5d %9.9s %.55s\n",
	// 		item.Number, item.User.Login, item.Title)
	// 	if time.Now().AddDate(-1, 0, 0).After(item.CreatedAt) {
	// 		fmt.Println("Created before a year ago")
	// 	}
	// }

	var report = template.Must(template.New("issuelist").Parse(issueshtml.Templ))
	report.Execute(os.Stdout, result)
}
