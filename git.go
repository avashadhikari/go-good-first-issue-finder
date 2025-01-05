package main

import (
	"strings"
)

func getSearchStrings(orgStrings []string) (res []string) {
	searchText := `is:issue is:open language:Go no:assignee state:open label:"help wanted" label:"good first issue" -is:archived`
	for _, s := range orgStrings {
		res = append(res, strings.TrimSpace(s+searchText))
	}
	return
}
