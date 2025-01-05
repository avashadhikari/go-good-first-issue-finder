package main

import (
	"log"
)

func main() {
	// Get top 300 org names and create search strings for them
	orgStrings := GetOrgStrings()
	searchStrings := GetSearchStrings(orgStrings)

	// Get issues for each search string and write them to a file
	issues := GetIssues(searchStrings)
	if err := WriteObjectsToFile("issues.txt", issues, func(i Issue) string {
		return i.ToString()
	}); err != nil {
		log.Fatalf("error writing objects to file: %s", err)
	}
}
