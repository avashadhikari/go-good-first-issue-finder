# gfifinder

gfifinder (aka Good First Issue Finder) is a package to find issues that are labelled as "good first issue" and other filters defined in `git.go` for the top 300 companies as defined in https://gitstar-ranking.com/organizations

# TODO/ next-in-line
1. make some filter criteria configurable via command line arguments

# Running it locally
Running the code is pretty simple, you just have to do a `go run .` since I'm not using any other external packages.   
Here's a sample response
```
❯ go run .
2025/01/05 15:58:16 Attempting to fetch issues from GitHub API
2025/01/05 15:58:16 Fetching issues for search string 1
2025/01/05 15:58:18 Fetching page 1 of issues
2025/01/05 15:58:21 Fetching page 2 of issues
2025/01/05 15:58:23 Fetching issues for search string 2
2025/01/05 15:58:25 Fetching page 1 of issues
2025/01/05 15:58:28 Fetching issues for search string 3
2025/01/05 15:58:30 Fetching page 1 of issues
2025/01/05 15:58:30 Fetched 222 issues from GitHub API
```