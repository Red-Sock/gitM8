package model

import (
	"strings"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

const refPrefix = "refs/heads/"

type PushPayload struct {
	Ref          string       `json:"ref"`
	Before       string       `json:"before"`
	After        string       `json:"after"`
	Repository   Repo         `json:"repository"`
	Pusher       UserShort    `json:"pusher"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
	Created      bool         `json:"created"`
	Deleted      bool         `json:"deleted"`
	Forced       bool         `json:"forced"`
	BaseRef      interface{}  `json:"base_ref"`
	Compare      string       `json:"compare"`
	Commits      []Commit     `json:"commits"`
	HeadCommit   Commit       `json:"head_commit"`
}

func (p *PushPayload) GetAction() domain.Action {
	return domain.ActionUnknown
}

func (p *PushPayload) GetCommitsAmount() int {
	return len(p.Commits)
}

func (p *PushPayload) GetEventType() domain.EventType {
	return domain.Push
}

func (p *PushPayload) GetProject() domain.Project {
	return p.Repository.ToDomain()
}

func (p *PushPayload) GetAuthor() domain.Author {
	return p.Sender.ToDomain()
}

func (p *PushPayload) GetSrcBranch() domain.Branch {
	branchName := strings.Replace(p.Ref, refPrefix, "", 1)
	return domain.Branch{
		Name: branchName,
		Link: p.Repository.HtmlUrl + "/tree/" + branchName,
	}
}

func (p *PushPayload) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{}
}

func (p *PushPayload) GetCommits() []domain.Commit {
	out := make([]domain.Commit, 0, len(p.Commits))
	for _, item := range p.Commits {
		out = append(out, domain.Commit{
			Author: item.Author.ToDomain(),
		})
	}
	return out
}
