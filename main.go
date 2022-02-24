package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"

	"github.com/crqra/go-action/pkg/action"
)

var patterns = []*regexp.Regexp{
	regexp.MustCompile("(?im)^breaking change:.*"),
	regexp.MustCompile(`(?i)^(\w+)(\(.*\))?!:.*`),
	regexp.MustCompile(`(?i)^feat(\(.*\))?:.*`),
	regexp.MustCompile(`(?i)^fix(\(.*\))?:.*`),
	regexp.MustCompile(`(?i)^docs(\(.*\))?:.*`),
	regexp.MustCompile(`(?i)^ci(\(.*\))?:.*`),
	regexp.MustCompile(`(?i)^build(\(.*\))?:.*`),
}

type PullRequestEvent struct {
	PullRequest struct {
		Title string `json:"title"`
	} `json:"pull_request"`
}

func (evt PullRequestEvent) String() string {
	return evt.PullRequest.Title
}

type PushEvent struct {
	Commits []struct {
		ID      string `json:"id"`
		Message string `json:"message"`
	} `json:"commits"`
}

type ConventionalCommitsAction struct{}

func (a *ConventionalCommitsAction) Run() error {
	switch action.Context.EventName {
	case "pull_request":
		var evt PullRequestEvent

		if err := parseEvent(&evt); err != nil {
			return err
		}

		return validatePullRequest(evt)

	case "push":
		var evt PushEvent

		if err := parseEvent(&evt); err != nil {
			return err
		}

		return validatePush(evt)
	}

	return nil
}

func main() {
	if err := action.Execute(&ConventionalCommitsAction{}); err != nil {
		action.SetFailed(err, map[string]string{})
	}
}

func validatePullRequest(evt PullRequestEvent) error {
	valid := false

	for _, p := range patterns {
		valid = p.MatchString(evt.PullRequest.Title)
	}

	if !valid {
		return fmt.Errorf("Pull Request title is not a valid Conventional Commit")
	}

	return nil
}

func validatePush(evt PushEvent) error {
	for _, c := range evt.Commits {
		valid := false

		for _, p := range patterns {
			valid = p.MatchString(c.Message)
		}

		if !valid {
			return fmt.Errorf("Commit %s is not a valid Conventional Commit", c.ID)
		}
	}

	return nil
}

func parseEvent(evt interface{}) error {
	data, err := ioutil.ReadFile(action.Context.EventPath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, evt)
}
