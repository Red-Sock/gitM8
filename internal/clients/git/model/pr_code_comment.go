package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type PullRequestCodeComment struct {
	Action string `json:"action"`
	Review struct {
		Id                int       `json:"id"`
		NodeId            string    `json:"node_id"`
		User              User      `json:"user"`
		Body              string    `json:"body"`
		CommitId          string    `json:"commit_id"`
		SubmittedAt       time.Time `json:"submitted_at"`
		State             string    `json:"state"`
		HtmlUrl           string    `json:"html_url"`
		PullRequestUrl    string    `json:"pull_request_url"`
		AuthorAssociation string    `json:"author_association"`
		Links             struct {
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			PullRequest struct {
				Href string `json:"href"`
			} `json:"pull_request"`
		} `json:"_links"`
	} `json:"review"`
	PullRequest struct {
		Url                string        `json:"url"`
		Id                 int           `json:"id"`
		NodeId             string        `json:"node_id"`
		HtmlUrl            string        `json:"html_url"`
		DiffUrl            string        `json:"diff_url"`
		PatchUrl           string        `json:"patch_url"`
		IssueUrl           string        `json:"issue_url"`
		Number             int           `json:"number"`
		State              string        `json:"state"`
		Locked             bool          `json:"locked"`
		Title              string        `json:"title"`
		User               User          `json:"user"`
		Body               string        `json:"body"`
		CreatedAt          time.Time     `json:"created_at"`
		UpdatedAt          time.Time     `json:"updated_at"`
		ClosedAt           interface{}   `json:"closed_at"`
		MergedAt           interface{}   `json:"merged_at"`
		MergeCommitSha     string        `json:"merge_commit_sha"`
		Assignee           interface{}   `json:"assignee"`
		Assignees          []interface{} `json:"assignees"`
		RequestedReviewers []interface{} `json:"requested_reviewers"`
		RequestedTeams     []interface{} `json:"requested_teams"`
		Labels             []interface{} `json:"labels"`
		Milestone          interface{}   `json:"milestone"`
		Draft              bool          `json:"draft"`
		CommitsUrl         string        `json:"commits_url"`
		ReviewCommentsUrl  string        `json:"review_comments_url"`
		ReviewCommentUrl   string        `json:"review_comment_url"`
		CommentsUrl        string        `json:"comments_url"`
		StatusesUrl        string        `json:"statuses_url"`
		Head               struct {
			Label string `json:"label"`
			Ref   string `json:"ref"`
			Sha   string `json:"sha"`
			User  User   `json:"user"`
			Repo  Repo   `json:"repo"`
		} `json:"head"`
		Base  Branch `json:"base"`
		Links struct {
			Self struct {
				Href string `json:"href"`
			} `json:"self"`
			Html struct {
				Href string `json:"href"`
			} `json:"html"`
			Issue struct {
				Href string `json:"href"`
			} `json:"issue"`
			Comments struct {
				Href string `json:"href"`
			} `json:"comments"`
			ReviewComments struct {
				Href string `json:"href"`
			} `json:"review_comments"`
			ReviewComment struct {
				Href string `json:"href"`
			} `json:"review_comment"`
			Commits struct {
				Href string `json:"href"`
			} `json:"commits"`
			Statuses struct {
				Href string `json:"href"`
			} `json:"statuses"`
		} `json:"_links"`
		AuthorAssociation string      `json:"author_association"`
		AutoMerge         interface{} `json:"auto_merge"`
		ActiveLockReason  interface{} `json:"active_lock_reason"`
	} `json:"pull_request"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
}

func (p *PullRequestCodeComment) GetRelease() domain.ReleasePayload {
	return domain.ReleasePayload{}
}

func (p *PullRequestCodeComment) GetWorkflow() domain.Workflow {
	return domain.Workflow{}
}

func (p *PullRequestCodeComment) GetAction() string {
	return p.Action
}

func (p *PullRequestCodeComment) GetProject() domain.Project {
	return p.Repository.ToDomain()
}

func (p *PullRequestCodeComment) GetAuthor() domain.Author {
	return p.Sender.ToDomain()
}

func (p *PullRequestCodeComment) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{
		StateStr: p.PullRequest.State,
		Name:     p.PullRequest.Title,
		Link:     p.PullRequest.HtmlUrl,
	}
}

func (p *PullRequestCodeComment) GetCommits() []domain.Commit {
	return nil
}

func (p *PullRequestCodeComment) GetCommitsAmount() int {
	return 0
}

func (p *PullRequestCodeComment) GetSrcBranch() domain.Branch {
	return domain.Branch{}
}

func (p *PullRequestCodeComment) GetEventType() domain.EventType {
	return domain.ReviewComment
}
