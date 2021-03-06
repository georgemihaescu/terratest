// Package git allows to interact with Git.
package git

import (
	"os/exec"
	"strings"

	"github.com/gruntwork-io/terratest/modules/testing"
)

// GetCurrentBranchName retrieves the current branch name or
// empty string in case of detached state.
func GetCurrentBranchName(t testing.TestingT) string {
	out, err := GetCurrentBranchNameE(t)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// GetCurrentBranchNameE retrieves the current branch name or
// empty string in case of detached state.
func GetCurrentBranchNameE(t testing.TestingT) (string, error) {
	cmd := exec.Command("git", "rev-parse", "--abbrev-ref", "HEAD")
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}

	name := strings.TrimSpace(string(bytes))
	if name == "HEAD" {
		return "", nil
	}

	return name, nil
}

// GetCurrentGitRef retrieves current branch name, lightweight (non-annotated) tag or
// if tag points to the commit exact tag value.
func GetCurrentGitRef(t testing.TestingT) string {
	out, err := GetCurrentGitRefE(t)
	if err != nil {
		t.Fatal(err)
	}
	return out
}

// GetCurrentGitRefE retrieves current branch name, lightweight (non-annotated) tag or
// if tag points to the commit exact tag value.
func GetCurrentGitRefE(t testing.TestingT) (string, error) {
	out, err := GetCurrentBranchNameE(t)

	if err != nil {
		return "", err
	}

	if out != "" {
		return out, nil
	}

	out, err = GetTagE(t)
	if err != nil {
		return "", err
	}
	return out, nil
}

// GetTagE retrieves lightweight (non-annotated) tag or if tag points
// to the commit exact tag value.
func GetTagE(t testing.TestingT) (string, error) {
	cmd := exec.Command("git", "describe", "--tags")
	bytes, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytes)), nil
}
