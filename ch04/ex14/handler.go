package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type (
	parameter struct {
		repogitry  string
		owner      string
		issueQuery []string
		userQuery  []string
	}
	Result struct {
		Issues     *IssueResponse
		Milestones *MilestoneResponse
		Users      *UserResponse
	}
)

func handler(w http.ResponseWriter, r *http.Request) {
	params := &parameter{}
	q := r.URL.RawQuery
	m, _ := url.ParseQuery(q)

	if v, ok := m["repo"]; ok {
		params.repogitry = v[0]
	}
	if v, ok := m["owner"]; ok {
		params.owner = v[0]
	}
	if v, ok := m["i"]; ok {
		params.issueQuery = v
	}
	if v, ok := m["u"]; ok {
		params.userQuery = v
	}

	result := Result{}
	issues, err := getIssues(params)
	if err != nil {
		log.Println(err)
	}
	milestones, err := getMilestones(params)
	if err != nil {
		log.Println(err)
	}
	users, err := getUsers(params)
	if err != nil {
		log.Println(err)
	}
	result.Issues = issues
	result.Milestones = milestones
	result.Users = users

	t := template.Must(template.ParseFiles("info.html.tpl"))
	if err := t.Execute(w, result); err != nil {
		log.Fatal(err)
	}
}
