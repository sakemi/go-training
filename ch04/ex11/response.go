package main

type GetIssueResponse struct {
	URL       string        `json:"url"`
	State     string        `json:"state"`
	Title     string        `json:"title"`
	Body      string        `json:"body"`
	Assignee  AssigneeData  `json:"assignee"`
	Milestone MilestoneData `json:"milestone"`
}

type GetIssuesResponse struct {
	Issues []GetIssueResponse
}

type AssigneeData struct {
	Login string `json:"login"`
}

type MilestoneData struct {
	Number int    `json:"number"`
	Title  string `json:"title"`
}
