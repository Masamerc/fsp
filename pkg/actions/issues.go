package actions

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/cli/go-gh/v2"
	"github.com/urfave/cli/v2"
)

type Issue struct {
	Title    string
	Body     string
	Repo     string
	Assignee string
}

func readIssues(path string) ([]Issue, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	if err != nil {
		log.Fatal(err)
	}

	// discard the header
	_, err = reader.Read()
	if err != nil {
		return nil, err
	}

	issues := []Issue{}

	// read the rest
	for {
		record, err := reader.Read()
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			return nil, err
		}

		issue := Issue{
			Title:    record[0],
			Body:     record[1],
			Repo:     record[2],
			Assignee: record[3],
		}
		issues = append(issues, issue)
	}

	return issues, nil
}

func createIssue(issue Issue, labels []string, project string) {

	args := []string{
		"issue",
		"create",
		"--title",
		issue.Title,
		"--body",
		issue.Body,
		"--repo",
		issue.Repo,
		"--assignee",
		issue.Assignee,
	}

	if len(labels) > 0 {
		for _, label := range labels {
			args = append(args, "--label", label)
		}
	}

	if project != "" {
		args = append(args, "--project", project)
	}

	resp, _, err := gh.Exec(args...)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("issue created: %s", resp.String())
	fmt.Printf("repo: %s\n", issue.Repo)
	fmt.Printf("title: %s\n\n", issue.Title)
}

func BulkCreateIssues(c *cli.Context) error {
	issues, err := readIssues(c.String("file"))
	if err != nil {
		fmt.Println("Error reading issues from file:", err)
		return err
	}

	labels := c.StringSlice("labels")
	project := c.String("project-id")

	for _, issue := range issues {
		createIssue(issue, labels, project)
	}
	return nil
}
