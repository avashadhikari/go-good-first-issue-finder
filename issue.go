package main

import "fmt"

type Issue struct {
	Url   string `json:"html_url"`
	Title string `json:"title"`
}

type IssueAPIResponse struct {
	Items []Issue `json:"items"`
}

func (i Issue) ToString() string {
	return fmt.Sprintf("%s: %s", i.Title, i.Url)
}
