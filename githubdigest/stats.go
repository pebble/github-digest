package githubdigest

import (
	"time"

	"github.com/pebble/github-digest/Godeps/_workspace/src/github.com/google/go-github/github"
)

type PullRequestStats struct {
	Project      string    `json:"project"`
	Number       int       `json:"number"`
	Title        string    `json:"title"`
	User         string    `json:"user"`
	MergedBy     *string   `json:"merged_by,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Additions    int       `json:"additions"`
	Deletions    int       `json:"deletions"`
	ChangedFiles int       `json:"changed_files"`
}

type UserStats struct {
	Open     uint32 `json:"open"`
	Closed   uint32 `json:"closed"`
	Comments uint32 `json:"comments"`
}

type GithubDigest struct {
	Open   []PullRequestStats    `json:"open"`
	Closed []PullRequestStats    `json:"closed"`
	Users  map[string]*UserStats `json:"users"`
	Repos  []string              `json:"repos"`
}

func NewGithubDigest(repositories []string) *GithubDigest {
	return &GithubDigest{
		Repos:  repositories,
		Users:  make(map[string]*UserStats),
		Open:   make([]PullRequestStats, 0),
		Closed: make([]PullRequestStats, 0),
	}
}

func NewPullRequestStats(project string, pull github.PullRequest) PullRequestStats {
	var updatedAt = *pull.UpdatedAt
	var mergedBy *string
	if pull.MergedBy != nil {
		mergedBy = &(*pull.MergedBy.Login)
		// Ignore post-merge comments:
		updatedAt = *pull.MergedAt
	}

	return PullRequestStats{
		Project:      project,
		Number:       *pull.Number,
		Title:        *pull.Title,
		User:         *pull.User.Login,
		MergedBy:     mergedBy,
		CreatedAt:    *pull.CreatedAt,
		UpdatedAt:    updatedAt,
		Additions:    *pull.Additions,
		Deletions:    *pull.Deletions,
		ChangedFiles: *pull.ChangedFiles,
	}
}

func (d GithubDigest) GetUser(user *github.User) *UserStats {
	username := *user.Login
	if stats, exists := d.Users[username]; exists {
		return stats
	} else {
		stats := &UserStats{
			Open:     0,
			Closed:   0,
			Comments: 0,
		}
		d.Users[username] = stats
		return stats
	}
}
