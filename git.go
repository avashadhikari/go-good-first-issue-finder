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
	"time"
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
	for i, s := range searchStrings {
		log.Printf("Fetching issues for search string %d\n", i+1)
		res := makeApiCall(s)
		issues = append(issues, res...)
	}
	log.Printf("Fetched %d issues from GitHub API\n", len(issues))
	return
}

// makeApiCall makes a call to the GitHub API and returns a slice of issues
func makeApiCall(query string) (issues []Issue) {
	url := "https://api.github.com/search/issues"
	token := os.Getenv("GITHUB_TOKEN")
	perPage := 100
	page := 1
	client := http.Client{}

	ticker := time.NewTicker(time.Minute / 30) // 30 API calls per minute
	defer ticker.Stop()

	for {
		<-ticker.C // Wait for the ticker
		log.Printf("Fetching page %d of issues\n", page)
		req := prepareRequest(url, query, token, page, perPage)
		res, err := client.Do(req)
		if err != nil {
			log.Fatalf("error making request: %s\n", err)
		}
		issues = append(issues, mapAPIResponseToIssues(res)...)

		// Check if there are more pages
		if res.Header.Get("Link") == "" || !strings.Contains(res.Header.Get("Link"), `rel="next"`) {
			break
		}
		page++
	}
	return
}

// prepareRequest prepares a http request with the necessary headers and query parameters
func prepareRequest(url, query, token string, page, perPage int) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, url, nil)
	req.Header = http.Header{
		"Accept":               {"application/vnd.github.v3+json"},
		"X-GitHub-Api-Version": {"2022-11-28"},
		"Authorization":        {fmt.Sprintf("Bearer %s", token)},
	}
	q := req.URL.Query()
	q.Add("q", query)
	q.Add("order", "desc")
	q.Add("per_page", strconv.Itoa(perPage))
	q.Add("page", strconv.Itoa(page))
	q.Add("sort", "updated")
	req.URL.RawQuery = q.Encode()
	return req
}

// mapAPIResponseToIssues maps the response from the GitHub API to a slice of Issue structs
func mapAPIResponseToIssues(res *http.Response) (issues []Issue) {
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
		log.Fatalf("Error unmarshalling response: %s\n", err)
	}
	return apiResponse.Items
}
