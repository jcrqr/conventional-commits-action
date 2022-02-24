package main

import "testing"

func TestValidatePullRequestOnValidTitle(t *testing.T) {
	evt := PullRequestEvent{
		PullRequest: PullRequest{
			Title: "fix: something",
		},
	}

	if err := validatePullRequest(evt); err != nil {
		t.Fail()
	}
}

func TestValidatePullRequestOnInvalidTitle(t *testing.T) {
	evt := PullRequestEvent{
		PullRequest: PullRequest{
			Title: "wrong pull request title",
		},
	}

	if err := validatePullRequest(evt); err == nil {
		t.Fail()
	}
}

func TestValidatePushOnValidCommits(t *testing.T) {
	evt := PushEvent{
		Commits: []Commit{
			{ID: "4204bd1", Message: "fix: something"},
			{ID: "4204bd3", Message: "feat: something"},
		},
	}

	if err := validatePush(evt); err != nil {
		t.Fail()
	}
}

func TestValidatePushOnInvalidCommits(t *testing.T) {
	evt := PushEvent{
		Commits: []Commit{
			{ID: "4204bd1", Message: "fix: something"},
			{ID: "4204bd2", Message: "wrong message"},
			{ID: "4204bd3", Message: "feat: something"},
		},
	}

	if err := validatePush(evt); err == nil {
		t.Fail()
	}
}
