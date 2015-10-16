package githubdigest

import "time"
import "strings"
import "golang.org/x/oauth2"
import (
	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/google/go-github/github"

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

func NewDigester(oauth_token string) GithubDigester {
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: oauth_token},
	)
	tc := oauth2.NewClient(oauth2.NoContext, ts)
	client := github.NewClient(tc)
	return GithubDigester{client: client}
}


func (d GithubDigester) GetDigest(repositories []string, stat_cutoff time.Time, closed_cutoff time.Time) *GithubDigest {
	stats := NewGithubDigest()

	openPr := &github.PullRequestListOptions{State: "open", Sort: "updated", Direction: "desc"}
	closedPr := &github.PullRequestListOptions{State: "closed", Sort: "updated", Direction: "desc"}

	for _, repoKey := range repositories {
		owner, repo := parseRepo(repoKey)

		// List open PRs:
		open_pulls, _, _ := d.client.PullRequests.List(owner, repo, openPr)
		for _, pull := range open_pulls {
			// User stat: PR opened
			stats.GetUser(pull.User).Open += 1

			// This is an open PR: store
			pull_detail, _, _ := d.client.PullRequests.Get(owner, repo, *pull.Number)
			stats.Open = append(stats.Open, NewPullRequestStats(repoKey, *pull_detail))

			d.addCommentStats(owner, repo, *pull.Number, stats)
		}

		// List closed PRs:
		closed_pulls, _, _ := d.client.PullRequests.List(owner, repo, closedPr)
		for _, pull := range closed_pulls {
			// Ignore PRs older than cutoff:
			if stat_cutoff.After(*pull.CreatedAt) {
				break
			}

			// Only track stats on PRs that were actually merged (not dropped)
			pullDetail, _, _ := d.client.PullRequests.Get(owner, repo, *pull.Number)
			if pullDetail.Merged != nil && *pullDetail.Merged {
				if closed_cutoff.Before(*pullDetail.MergedAt) {
					// This is within ClosedCutoff, store
					stats.Closed = append(stats.Closed, NewPullRequestStats(repoKey, *pullDetail))
				}

				// User stat: PR opened
				stats.GetUser(pull.User).Open += 1

				// User stat: PR merged
				stats.GetUser(pullDetail.MergedBy).Closed += 1

				d.addCommentStats(owner, repo, *pull.Number, stats)
			}
		}
	}

	return stats
}


func (d GithubDigester) addCommentStats(owner string, repo string, pull int, stats *GithubDigest) {
	prComments, _, _ := d.client.PullRequests.ListComments(owner, repo, pull, nil)
	for _, comment := range prComments {
		// User stat: comment made
		stats.GetUser(comment.User).Comments += 1
	}

	issueComments, _, _ := d.client.Issues.ListComments(owner, repo, pull, nil)
	for _, comment := range issueComments {
		// User stat: comment made
		stats.GetUser(comment.User).Comments += 1
	}
}