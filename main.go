package main

import (
	"fmt"
	"regexp"

	"github.com/crqra/go-action/pkg/action"
	"github.com/google/go-github/v42/github"
)

var pattern = regexp.MustCompile(`(?i)^(\w+)(\(.*\))?:.*`)

type ConventionalCommitsAction struct{}

func (a *ConventionalCommitsAction) Run() error {
	evt, err := action.GetEvent()
	if err != nil {
		return err
	}

	switch evt := evt.(type) {
	case *github.PullRequestEvent:
		if match := pattern.MatchString(evt.PullRequest.GetTitle()); !match {
			return fmt.Errorf("Pull Request title is not a valid Conventional Commit")
		}

		return nil

	case *github.PushEvent:
		for _, c := range evt.Commits {
			if match := pattern.MatchString(c.GetMessage()); !match {
				return fmt.Errorf("Commit %s is not a valid Conventional Commit", c.GetSHA())
			}
		}

		return nil

	default:
		action.Notice(
			"conventional-commits-action skipped: only runs for pull_request and push events",
			map[string]string{},
		)
	}

	return nil
}

func main() {
	if err := action.Execute(&ConventionalCommitsAction{}); err != nil {
		action.SetFailed(err, map[string]string{})
	}
}
