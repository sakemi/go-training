package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const api = "https://api.github.com"

func createIssue() error {
	var l []string
	var owner, repo, mode string
	var labels string
	var createReq = new(CreateRequest)

	fmt.Print("Input owner:")
	fmt.Scan(&owner)
	fmt.Print("Input repogitry:")
	fmt.Scan(&repo)
	fmt.Print("Input title:")
	fmt.Scan(&createReq.Title)

	fmt.Print("Run editor to edit Body (y or n):")
	for {
		fmt.Scan(&mode)
		if mode == "y" || mode == "n" {
			break
		}
		fmt.Print("Input y or n:")
	}
	if mode == "y" {
		var editor string
		fmt.Print("Input editor to run:")
		fmt.Scan(&editor)
		body, err := runEditor(editor)
		if err != nil {
			return err
		}
		createReq.Body = body
	} else {
		fmt.Print("Input body:")
		fmt.Scan(&createReq.Body)
	}

	fmt.Print("Input assignee:")
	fmt.Scan(&createReq.Assignee)
	fmt.Print("Input milestone:")
	fmt.Scan(&createReq.Milestone)
	fmt.Print("Input labels devided by space:")
	fmt.Scan(&labels)
	createReq.Labels = strings.Split(labels, " ")

	l = append(l, api, "repos", owner, repo, "issues")
	issueURL := strings.Join(l, "/")
	params, err := json.Marshal(createReq)
	if err != nil {
		return err
	}

	return postJSON(issueURL, params)
}

func readIssue() error {
	var owner, repo, mode, number string
	var l []string

	fmt.Print("Input owner:")
	fmt.Scan(&owner)
	fmt.Print("Input repogitry:")
	fmt.Scan(&repo)
	fmt.Print("Specify issue number? (y or n):")
	for {
		fmt.Scan(&mode)
		if mode == "y" || mode == "n" {
			break
		}
		fmt.Print("Input y or n:")
	}

	// single issue
	if mode == "y" {
		fmt.Print("Input issue number if you want to specify by issue number:")
		fmt.Scan(&number)

		l = append(l, api, "repos", owner, repo, "issues", number)
		issueURL := strings.Join(l, "/")
		resp, err := http.Get(issueURL)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		var data GetIssueResponse
		if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
			return err
		}
		showIssue(&data)
		return nil
	}

	//list issue
	var milestone, state, assignee, creator, labels string
	params := make(map[string]string)
	fmt.Print("Input milestone:")
	fmt.Scan(&milestone)
	params["milestone"] = milestone
	fmt.Print("Input state:")
	fmt.Scan(&state)
	params["state"] = state
	fmt.Print("Input assignee:")
	fmt.Scan(&assignee)
	params["assignee"] = assignee
	fmt.Print("Input creator:")
	fmt.Scan(&creator)
	params["creator"] = creator
	fmt.Print("Input labels:")
	fmt.Scan(&labels)
	params["labels"] = labels

	param := []string{"/?"}
	for k, v := range params {
		param = append(param, url.QueryEscape(k), "=", url.QueryEscape(v), "&")
	}
	l = append(l, api, "repos", owner, repo, "issues")
	var p []string
	p = append(p, param[:len(param)-1]...)
	issueURL := strings.Join(l, "/") + strings.Join(p, "")
	fmt.Println(issueURL)
	resp, err := http.Get(issueURL)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	var data GetIssuesResponse
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		return err
	}
	for _, v := range data.Issues {
		showIssue(&v)
	}
	return nil
}

func showIssue(issue *GetIssueResponse) {
	fmt.Println("====================")
	fmt.Printf("URL:%s\n", issue.URL)
	fmt.Printf("State:%s\n", issue.State)
	fmt.Printf("Title:%s\n", issue.Title)
	fmt.Printf("Body:%s\n", issue.Body)
	fmt.Printf("Assignee:%s\n", issue.Assignee.Login)
	fmt.Println("Milestone")
	fmt.Printf("\tNumber:%d\n", issue.Milestone.Number)
	fmt.Printf("\tTitle:%s\n", issue.Milestone.Title)
	fmt.Println("====================")
}

func updateIssue() error {
	var l []string
	var owner, repo, number, mode string
	var labels string
	var updateReq = new(UpdateRequest)

	fmt.Print("Input owner:")
	fmt.Scan(&owner)
	fmt.Print("Input repogitry:")
	fmt.Scan(&repo)
	fmt.Print("Input issue number:")
	fmt.Scan(&number)
	fmt.Print("Input title:")
	fmt.Scan(&updateReq.Title)

	fmt.Print("Run editor to edit Body (y or n):")
	for {
		fmt.Scan(&mode)
		if mode == "y" || mode == "n" {
			break
		}
		fmt.Print("Input y or n:")
	}
	if mode == "y" {
		var editor string
		fmt.Print("Input editor to run:")
		fmt.Scan(&editor)
		body, err := runEditor(editor)
		if err != nil {
			return err
		}
		updateReq.Body = body
	} else {
		fmt.Print("Input body:")
		fmt.Scan(&updateReq.Body)
	}

	fmt.Print("Input state(open or closed):")
	fmt.Scan(&updateReq.State)
	fmt.Print("Input assignee:")
	fmt.Scan(&updateReq.Assignee)
	fmt.Print("Input milestone:")
	fmt.Scan(&updateReq.Milestone)
	fmt.Print("Input labels devided by space:")
	fmt.Scan(&labels)
	updateReq.Labels = strings.Split(labels, " ")

	l = append(l, api, "repos", owner, repo, "issues", number)
	issueURL := strings.Join(l, "/")
	params, err := json.Marshal(updateReq)
	if err != nil {
		return err
	}

	return postJSON(issueURL, params)
}

func postJSON(url string, json []byte) error {
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json))
	if err != nil {
		return err
	}
	setHeaders(req)
	if err := setAuthToken(req); err != nil {
		return fmt.Errorf("Failed to read token: %v", err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	if (resp.StatusCode != http.StatusCreated) && (resp.StatusCode != http.StatusOK) {
		b, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(b))
		resp.Body.Close()
		return fmt.Errorf("Failed to create: %s", resp.Status)
	}
	resp.Body.Close()
	fmt.Println("Success!")
	return nil
}

func setHeaders(req *http.Request) {
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "sakemi")
}

func setAuthToken(req *http.Request) error {
	data, err := ioutil.ReadFile("token.txt")
	if err != nil {
		return err
	}
	token := strings.TrimRight(string(data), "\n")
	token = strings.TrimRight(token, "\r")
	req.SetBasicAuth("sakemi", token)
	return nil
}
