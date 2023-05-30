package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type PullRequestCodeComment struct {
	Action  string `json:"action"`
	Comment struct {
		Url                 string `json:"url"`
		PullRequestReviewId int    `json:"pull_request_review_id"`
		Id                  int    `json:"id"`
		NodeId              string `json:"node_id"`
		DiffHunk            string `json:"diff_hunk"`
		Path                string `json:"path"`
		CommitId            string `json:"commit_id"`
		OriginalCommitId    string `json:"original_commit_id"`
		User                struct {
			Login             string `json:"login"`
			Id                int    `json:"id"`
			NodeId            string `json:"node_id"`
			AvatarUrl         string `json:"avatar_url"`
			GravatarId        string `json:"gravatar_id"`
			Url               string `json:"url"`
			HtmlUrl           string `json:"html_url"`
			FollowersUrl      string `json:"followers_url"`
			FollowingUrl      string `json:"following_url"`
			GistsUrl          string `json:"gists_url"`
			StarredUrl        string `json:"starred_url"`
			SubscriptionsUrl  string `json:"subscriptions_url"`
			OrganizationsUrl  string `json:"organizations_url"`
			ReposUrl          string `json:"repos_url"`
			EventsUrl         string `json:"events_url"`
			ReceivedEventsUrl string `json:"received_events_url"`
			Type              string `json:"type"`
			SiteAdmin         bool   `json:"site_admin"`
		} `json:"user"`
		Body              string    `json:"body"`
		CreatedAt         time.Time `json:"created_at"`
		UpdatedAt         time.Time `json:"updated_at"`
		HtmlUrl           string    `json:"html_url"`
		PullRequestUrl    string    `json:"pull_request_url"`
		AuthorAssociation string    `json:"author_association"`
		Links             Links     `json:"_links"`
		Reactions         struct {
			Url        string `json:"url"`
			TotalCount int    `json:"total_count"`
			Field3     int    `json:"+1"`
			Field4     int    `json:"-1"`
			Laugh      int    `json:"laugh"`
			Hooray     int    `json:"hooray"`
			Confused   int    `json:"confused"`
			Heart      int    `json:"heart"`
			Rocket     int    `json:"rocket"`
			Eyes       int    `json:"eyes"`
		} `json:"reactions"`
		StartLine         int    `json:"start_line"`
		OriginalStartLine int    `json:"original_start_line"`
		StartSide         string `json:"start_side"`
		Line              int    `json:"line"`
		OriginalLine      int    `json:"original_line"`
		Side              string `json:"side"`
		OriginalPosition  int    `json:"original_position"`
		Position          int    `json:"position"`
		SubjectType       string `json:"subject_type"`
	} `json:"comment"`
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
		Body               interface{}   `json:"body"`
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
		Head               Branch        `json:"head"`
		Base               Branch        `json:"base"`
		Links              Link          `json:"_links"`
		AuthorAssociation  string        `json:"author_association"`
		AutoMerge          interface{}   `json:"auto_merge"`
		ActiveLockReason   interface{}   `json:"active_lock_reason"`
	} `json:"pull_request"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
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
	return domain.Comment
}
