package model

import (
	"strings"
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type PullRequestPayload struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		Url                 string        `json:"url"`
		Id                  int           `json:"id"`
		NodeId              string        `json:"node_id"`
		HtmlUrl             string        `json:"html_url"`
		DiffUrl             string        `json:"diff_url"`
		PatchUrl            string        `json:"patch_url"`
		IssueUrl            string        `json:"issue_url"`
		Number              int           `json:"number"`
		State               string        `json:"state"`
		Locked              bool          `json:"locked"`
		Title               string        `json:"title"`
		User                User          `json:"user"`
		Body                interface{}   `json:"body"`
		CreatedAt           time.Time     `json:"created_at"`
		UpdatedAt           time.Time     `json:"updated_at"`
		ClosedAt            interface{}   `json:"closed_at"`
		MergedAt            interface{}   `json:"merged_at"`
		MergeCommitSha      interface{}   `json:"merge_commit_sha"`
		Assignee            interface{}   `json:"assignee"`
		Assignees           []interface{} `json:"assignees"`
		RequestedReviewers  []interface{} `json:"requested_reviewers"`
		RequestedTeams      []interface{} `json:"requested_teams"`
		Labels              []interface{} `json:"labels"`
		Milestone           interface{}   `json:"milestone"`
		Draft               bool          `json:"draft"`
		CommitsUrl          string        `json:"commits_url"`
		ReviewCommentsUrl   string        `json:"review_comments_url"`
		ReviewCommentUrl    string        `json:"review_comment_url"`
		CommentsUrl         string        `json:"comments_url"`
		StatusesUrl         string        `json:"statuses_url"`
		Head                Branch        `json:"head"`
		Base                Branch        `json:"base"`
		Links               Links         `json:"_links"`
		AuthorAssociation   string        `json:"author_association"`
		AutoMerge           interface{}   `json:"auto_merge"`
		ActiveLockReason    interface{}   `json:"active_lock_reason"`
		Merged              bool          `json:"merged"`
		Mergeable           interface{}   `json:"mergeable"`
		Rebaseable          interface{}   `json:"rebaseable"`
		MergeableState      string        `json:"mergeable_state"`
		MergedBy            interface{}   `json:"merged_by"`
		Comments            int           `json:"comments"`
		ReviewComments      int           `json:"review_comments"`
		MaintainerCanModify bool          `json:"maintainer_can_modify"`
		Commits             int           `json:"commits"`
		Additions           int           `json:"additions"`
		Deletions           int           `json:"deletions"`
		ChangedFiles        int           `json:"changed_files"`
	} `json:"pull_request"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
}

func (p *PullRequestPayload) GetAction() domain.Action {
	return domain.Action(p.Action)
}

func (p *PullRequestPayload) GetProject() domain.Project {
	return p.Repository.ToDomain()
}

func (p *PullRequestPayload) GetAuthor() domain.Author {
	return p.Sender.ToDomain()
}

func (p *PullRequestPayload) GetSrcBranch() domain.Branch {
	return p.PullRequest.Base.ToDomain(p.Repository.HtmlUrl)
}

func (p *PullRequestPayload) GetPullRequest() domain.PullRequestPayload {
	pr := domain.PullRequestPayload{
		StateStr: p.PullRequest.State,
		Name:     p.PullRequest.Title,
		Link:     p.PullRequest.HtmlUrl,
		Base: domain.Branch{
			Name: p.PullRequest.Head.Ref,
			Link: strings.Join([]string{p.PullRequest.Head.Repo.HtmlUrl, p.PullRequest.Head.Ref}, ","),
		},
		Target: domain.Branch{
			Name: p.PullRequest.Head.Ref,
			Link: strings.Join([]string{p.PullRequest.Base.Repo.HtmlUrl, p.PullRequest.Base.Ref}, ","),
		},
	}

	return pr
}

func (p *PullRequestPayload) GetCommits() []domain.Commit {
	return nil
}

func (p *PullRequestPayload) GetCommitsAmount() int {
	return p.PullRequest.Commits
}

func (p *PullRequestPayload) GetEventType() domain.EventType {
	return domain.PullRequest
}

func (p *PullRequestPayload) GetWorkflow() domain.Workflow {
	return domain.Workflow{}
}
