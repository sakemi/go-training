package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const api = "https://api.github.com"

type (
	IssueResponse struct {
		TotalCount int `json:"total_count"`
		Items      []*Issue
	}
	Issue struct {
		Number    int
		HTMLURL   string `json:"html_url"`
		Title     string
		State     string
		User      *User
		CreatedAt time.Time `json:"created_at"`
		Body      string
	}
	User struct {
		Login   string
		HTMLURL string `json:"html_url"`
	}
	MilestoneResponse struct {
		Milestones []Milestone
	}
	Milestone struct {
		Number int
		State  string
		Title  string
	}
	UserResponse struct {
		TotalCount int `json:"total_count"`
		Items      []*User
	}
)

func getIssues(params *parameter) (*IssueResponse, error) {
	q := url.QueryEscape(strings.Join(params.issueQuery, " "))
	url := []string{api, "search", "issues"}
	issueURL := strings.Join(url, "/")
	if len(q) != 0 {
		issueURL += "?q=" + q
	}
	log.Println(issueURL)
	resp, err := http.Get(issueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to search issues: %s", resp.Status)
	}
	var result IssueResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func getMilestones(params *parameter) (*MilestoneResponse, error) {
	url := []string{api, "repos", params.owner, params.repogitry, "milestones"}
	issueURL := strings.Join(url, "/")
	log.Println(issueURL)
	resp, err := http.Get(issueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to get milestone: %s", resp.Status)
	}
	milestones := make([]Milestone, 0)
	if err := json.NewDecoder(resp.Body).Decode(&milestones); err != nil {
		return nil, err
	}
	result := MilestoneResponse{milestones}
	return &result, nil
}

func getUsers(params *parameter) (*UserResponse, error) {
	q := url.QueryEscape(strings.Join(params.userQuery, " "))
	url := []string{api, "search", "users"}
	issueURL := strings.Join(url, "/")
	if len(q) != 0 {
		issueURL += "?q=" + q
	}
	log.Println(issueURL)
	resp, err := http.Get(issueURL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Failed to search users: %s", resp.Status)
	}
	var result UserResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}
