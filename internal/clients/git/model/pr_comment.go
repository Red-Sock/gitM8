package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type PullRequestComment struct {
	Action string `json:"action"`
	Issue  struct {
		Url               string        `json:"url"`
		RepositoryUrl     string        `json:"repository_url"`
		LabelsUrl         string        `json:"labels_url"`
		CommentsUrl       string        `json:"comments_url"`
		EventsUrl         string        `json:"events_url"`
		HtmlUrl           string        `json:"html_url"`
		Id                int           `json:"id"`
		NodeId            string        `json:"node_id"`
		Number            int           `json:"number"`
		Title             string        `json:"title"`
		User              User          `json:"user"`
		Labels            []interface{} `json:"labels"`
		State             string        `json:"state"`
		Locked            bool          `json:"locked"`
		Assignee          interface{}   `json:"assignee"`
		Assignees         []interface{} `json:"assignees"`
		Milestone         interface{}   `json:"milestone"`
		Comments          int           `json:"comments"`
		CreatedAt         time.Time     `json:"created_at"`
		UpdatedAt         time.Time     `json:"updated_at"`
		ClosedAt          interface{}   `json:"closed_at"`
		AuthorAssociation string        `json:"author_association"`
		ActiveLockReason  interface{}   `json:"active_lock_reason"`
		Draft             bool          `json:"draft"`
		PullRequest       struct {
			Url      string      `json:"url"`
			HtmlUrl  string      `json:"html_url"`
			DiffUrl  string      `json:"diff_url"`
			PatchUrl string      `json:"patch_url"`
			MergedAt interface{} `json:"merged_at"`
		} `json:"pull_request"`
		Body                  interface{} `json:"body"`
		Reactions             Reaction    `json:"reactions"`
		TimelineUrl           string      `json:"timeline_url"`
		PerformedViaGithubApp interface{} `json:"performed_via_github_app"`
		StateReason           interface{} `json:"state_reason"`
	} `json:"issue"`
	Comment      Comment      `json:"comment"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
}

func (p *PullRequestComment) GetAction() domain.Action {
	return domain.Action(p.Action)
}

func (p *PullRequestComment) GetProject() domain.Project {
	return p.Repository.ToDomain()
}

func (p *PullRequestComment) GetAuthor() domain.Author {
	return p.Sender.ToDomain()
}

func (p *PullRequestComment) GetSrcBranch() domain.Branch {
	return domain.Branch{}
}

func (p *PullRequestComment) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{
		StateStr: p.Issue.State,
		Name:     p.Issue.Title,
		Link:     p.Issue.HtmlUrl,
	}
}

func (p *PullRequestComment) GetCommits() []domain.Commit {
	return nil
}

func (p *PullRequestComment) GetCommitsAmount() int {
	return 0
}

func (p *PullRequestComment) GetEventType() domain.EventType {
	return domain.IssueComment
}
