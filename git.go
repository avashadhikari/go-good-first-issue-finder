package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

// GetSearchStrings returns a slice of strings that will be used to search for issues on GitHub
// TODO: make search filter configurable via command line arguments
func GetSearchStrings(orgStrings []string) (res []string) {
	searchText := ` is:issue is:open language:Go no:assignee state:open label:"help wanted" label:"good first issue" -is:archived`
	for _, s := range orgStrings {
		res = append(res, strings.TrimSpace(s+searchText))
	}
	return
}

// GetIssues returns a slice of issues from the GitHub API
func GetIssues(searchStrings []string) (issues []Issue) {
	log.Println("Attempting to fetch issues from GitHub API")
	for _, s := range searchStrings {
		res := makeApiCall(s)
		issues = append(issues, res...)
	}
	log.Printf("Fetched %d issues from GitHub API\n", len(issues))
	return
}

// makeApiCall makes a call to the GitHub API and returns a slice of issues
// TODO: handle pagination, and limit api calls made to 30
func makeApiCall(query string) (issues []Issue) {
	url := "https://api.github.com/search/issues"
	token := os.Getenv("GITHUB_TOKEN")
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = http.Header{
		"Accept":               {"application/vnd.github.v3+json"},
		"X-GitHub-Api-Version": {"2022-11-28"},
		"Authorization":        {fmt.Sprintf("Bearer %s", token)},
	}

	perPage := 100
	page := 1
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("order", "desc")
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("page", strconv.Itoa(page))
	req.URL.RawQuery = q.Encode()

	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatalf("error making request: %s\n", err)
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s\n", err)
	}
	if res.StatusCode != http.StatusOK {
		log.Fatalln("status code is not ok: ", res.StatusCode)
	}
	var apiResponse IssueAPIResponse
	err = json.Unmarshal(body, &apiResponse)
	if err != nil {
		log.Fatalf("Error unmarshalling the response body: %s\n", err)
	}
	return apiResponse.Items
}
