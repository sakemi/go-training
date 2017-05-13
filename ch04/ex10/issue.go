package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	var monthly, yearly, overYear []*github.Issue
	result, err := github.SearchIssues(os.Args[1:])

	now := time.Now()
	for _, issue := range result.Items {
		if issue.CreatedAt.Before(now.AddDate(-1, 0, 0)) {
			overYear = append(overYear, issue)
		} else if issue.CreatedAt.Before(now.AddDate(0, -1, 0)) {
			yearly = append(yearly, issue)
		} else {
			monthly = append(monthly, issue)
		}
	}

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	fmt.Println("===1ヶ月未満===")
	for _, item := range monthly {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("===1年未満===")
	for _, item := range yearly {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
	fmt.Println("===それ以前===")
	for _, item := range overYear {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
	}
}
