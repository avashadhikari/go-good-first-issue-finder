package main

import "fmt"

func main() {
	orgStrings := getOrgStrings()
	searchStrings := getSearchStrings(orgStrings)
	fmt.Println(searchStrings)

}
