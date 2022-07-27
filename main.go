package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

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
	app := newApp(
		newGitHubClient(ctx, token),
		&pullRequest{
			owner:  os.Getenv("INPUT_OWNER"),
			repo:   os.Getenv("INPUT_REPO"),
			number: prNum,
		},
	)
	err = app.run(ctx)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

type app struct {
	githubClient *github.Client
	pr           *pullRequest
}

type pullRequest struct {
	owner  string
	repo   string
	number int
}

func newApp(githubClient *github.Client, pr *pullRequest) *app {
	return &app{
		githubClient: githubClient,
		pr:           pr,
	}
}

func newGitHubClient(ctx context.Context, token string) *github.Client {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return github.NewClient(tc)
}

func (a *app) run(ctx context.Context) error {
	ic, _, err := a.githubClient.Issues.CreateComment(ctx, a.pr.owner, a.pr.repo, a.pr.number, &github.IssueComment{Body: github.String("Deployed")})
	if err != nil {
		return err
	}
	fmt.Printf("::set-output name=url::%s\n", *ic.URL)
	return nil
}
