package ghrm

import (
	"context"

	"github.com/google/go-github/v32/github"
	"golang.org/x/oauth2"
)

// NewGitHubClient returns a GitHub client.
func NewGitHubClient(token string) *github.Client {
	ctx := context.Background()
	sts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return github.NewClient(oauth2.NewClient(ctx, sts))
}

// RemoveRepository removes a GitHub repository.
func RemoveRepository(gh *github.Client, owner, repo string) error {
	ctx := context.Background()
	if _, err := gh.Repositories.Delete(ctx, owner, repo); err != nil {
		return err
	}
	return nil
}
