package main

import (
	"testing"

	"github.com/crqra/go-action/pkg/action"
)

func TestValidPullRequest(t *testing.T) {
	action.Context.EventName = "pull_request"
	action.Context.EventPath = "testdata/valid_pr_event.json"

	if err := action.Execute(&ConventionalCommitsAction{}); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidPullRequest(t *testing.T) {
	action.Context.EventName = "pull_request"
	action.Context.EventPath = "testdata/invalid_pr_event.json"

	if err := action.Execute(&ConventionalCommitsAction{}); err == nil {
		t.Fatal(err)
	}
}

func TestValidPush(t *testing.T) {
	action.Context.EventName = "push"
	action.Context.EventPath = "testdata/valid_push_event.json"

	if err := action.Execute(&ConventionalCommitsAction{}); err != nil {
		t.Fatal(err)
	}
}

func TestInvalidPush(t *testing.T) {
	action.Context.EventName = "push"
	action.Context.EventPath = "testdata/invalid_push_event.json"

	if err := action.Execute(&ConventionalCommitsAction{}); err == nil {
		t.Fatal(err)
	}
}
