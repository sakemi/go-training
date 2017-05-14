package main

type CreateRequest struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignee  string   `json:"assignee"`
	Milestone int      `json:"milestone"`
	Labels    []string `json:"labels"`
}

type UpdateRequest struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Assignee  string   `json:"assignee"`
	State     string   `json:"state"`
	Milestone int      `json:"milestone"`
	Labels    []string `json:"labels"`
}
