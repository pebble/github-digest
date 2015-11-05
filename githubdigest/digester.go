package githubdigest

import (
	"strings"
	"time"

	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/google/go-github/github"
	"github.com/pebble/github-digest/Godeps/_workspace/src/golang.org/x/oauth2"
)

type GithubDigester struct {
	client *github.Client
}

func parseRepo(repo string) (string, string) {
	split := strings.Split(repo, "/")
	if len(split) == 2 {
		return split[0], split[1]
	} else {
		return "pebble", repo
	}
}

func NewDigester(oauthToken string) GithubDigester {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauthToken},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return GithubDigester{client: client}
}

func (d GithubDigester) GetDigest(repositories []string, statCutoff time.Time, closedCutoff time.Time) (*GithubDigest, error) {
	stats := NewGithubDigest(repositories)

	openPr := &github.PullRequestListOptions{State: "open", Sort: "updated", Direction: "desc"}
	closedPr := &github.PullRequestListOptions{State: "closed", Sort: "updated", Direction: "desc"}

	for _, repoKey := range repositories {
		owner, repo := parseRepo(repoKey)

		// List open PRs:
		openPulls, _, err := d.client.PullRequests.List(owner, repo, openPr)
		if err != nil {
			return nil, err
		}
		for _, pull := range openPulls {
			// User stat: PR opened
			stats.GetUser(pull.User).Open += 1

			// This is an open PR: store
			pullDetail, _, err := d.client.PullRequests.Get(owner, repo, *pull.Number)
			if err != nil {
				return nil, err
			}
			stats.Open = append(stats.Open, NewPullRequestStats(repoKey, *pullDetail))

			err = d.addCommentStats(owner, repo, *pull.Number, stats)
			if err != nil {
				return nil, err
			}

		}

		// List closed PRs:
		closedPulls, _, err := d.client.PullRequests.List(owner, repo, closedPr)
		if err != nil {
			return nil, err
		}
		for _, pull := range closedPulls {
			// Ignore PRs older than cutoff:
			if statCutoff.After(*pull.CreatedAt) {
				break
			}

			// Only track stats on PRs that were actually merged (not dropped)
			pullDetail, _, err := d.client.PullRequests.Get(owner, repo, *pull.Number)
			if err != nil {
				return nil, err
			}
			if pullDetail.Merged != nil && *pullDetail.Merged {
				if closedCutoff.Before(*pullDetail.MergedAt) {
					// This is within ClosedCutoff, store
					stats.Closed = append(stats.Closed, NewPullRequestStats(repoKey, *pullDetail))
				}

				// User stat: PR opened
				stats.GetUser(pull.User).Open += 1

				// User stat: PR merged
				stats.GetUser(pullDetail.MergedBy).Closed += 1

				err = d.addCommentStats(owner, repo, *pull.Number, stats)
				if err != nil {
					return nil, err
				}
			}
		}
	}

	return stats, nil
}

func (d GithubDigester) addCommentStats(owner string, repo string, pull int, stats *GithubDigest) error {
	prComments, _, err := d.client.PullRequests.ListComments(owner, repo, pull, nil)
	if err != nil {
		return err
	}
	for _, comment := range prComments {
		// User stat: comment made
		stats.GetUser(comment.User).Comments += 1
	}

	issueComments, _, err := d.client.Issues.ListComments(owner, repo, pull, nil)
	if err != nil {
		return err
	}
	for _, comment := range issueComments {
		// User stat: comment made
		stats.GetUser(comment.User).Comments += 1
	}
	return nil
}
