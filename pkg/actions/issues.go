package actions

import (
	"embed"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/cli/go-gh/v2"
	"github.com/urfave/cli/v2"
)

//go:embed templates/issue-template.md
var content embed.FS

func createTempBodyFile(body string) (string, error) {

	file, err := content.Open("templates/issue-template.md")
	if err != nil {
		return "", err
	}

	defer file.Close()

	issueTemplate, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}

	issueBody := strings.Replace(string(issueTemplate), "%BODY%", body, 1)

	tempFile, err := os.CreateTemp("", "issue-body")
	if err != nil {
		return "", err
	}

	_, err = tempFile.Write([]byte(issueBody))
	if err != nil {
		return "", err
	}

	tempFile.Close()
	fileName := tempFile.Name()
	return fileName, nil
}

type Issue struct {
	Title     string
	Body      string
	Repo      string
	Assignee  string
	Milestone string
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
			Title:     record[0],
			Body:      record[1],
			Repo:      record[2],
			Assignee:  record[3],
			Milestone: record[4],
		}
		issues = append(issues, issue)
	}

	return issues, nil
}

func createIssue(issue Issue, labels []string, project string) {

	bodyFile, err := createTempBodyFile(issue.Body)
	if err != nil {
		log.Fatal(err)
	}

	args := []string{
		"issue",
		"create",
		"--title",
		issue.Title,
		"--body-file",
		bodyFile,
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

	if issue.Milestone != "" {
		args = append(args, "--milestone", issue.Milestone)
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
