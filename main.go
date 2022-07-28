package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"text/template"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func main() {
	ctx := context.Background()
	token := os.Getenv("GITHUB_TOKEN")
	prNum, err := strconv.Atoi(os.Getenv("INPUT_PR_NUMBER"))
	if err != nil {
		log.Fatal(err)
	}
	pr := &pullRequest{
		owner:  os.Getenv("INPUT_OWNER"),
		repo:   os.Getenv("INPUT_REPO"),
		number: prNum,
	}
	comment, err := CostructComment(os.Getenv("INPUT_TEMPLATE"), os.Getenv("INPUT_FILEPATH"))
	if err != nil {
		log.Fatal(err)
	}
	gitHub := newGitHub(
		newGitHubClient(ctx, token),
	)
	url, err := gitHub.CreateComment(ctx, pr, comment)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("::set-output name=url::%s\n", url)
	os.Exit(0)
}

type GitHub struct {
	client  *github.Client
	pr      *pullRequest
	comment string
}

type pullRequest struct {
	owner  string
	repo   string
	number int
}

func newGitHub(client *github.Client) *GitHub {
	return &GitHub{
		client: client,
	}
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func (a *GitHub) CreateComment(ctx context.Context, pr *pullRequest, comment string) (string, error) {
	ic, _, err := a.client.Issues.CreateComment(
		ctx,
		a.pr.owner,
		a.pr.repo,
		a.pr.number,
		&github.IssueComment{Body: github.String(a.comment)},
	)
	if err != nil {
		return "", err
	}
	return *ic.URL, nil
}

func CostructComment(tmpl, filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	content := string(data)

	if tmpl == "" {
		tmpl = "{{.}}"
	}
	t, err := template.New("Comment template").Parse(tmpl)
	if err != nil {
		return "", err
	}
	var buf bytes.Buffer
	err = t.Execute(&buf, content)
	if err != nil {
		return "", err
	}
	return buf.String(), nil
}
