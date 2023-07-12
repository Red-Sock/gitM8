package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type Release struct {
	Action  string `json:"action"`
	Release struct {
		Url       string `json:"url"`
		AssetsUrl string `json:"assets_url"`
		UploadUrl string `json:"upload_url"`
		HtmlUrl   string `json:"html_url"`
		Id        int    `json:"id"`
		Author    struct {
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
		} `json:"author"`
		NodeId          string        `json:"node_id"`
		TagName         string        `json:"tag_name"`
		TargetCommitish string        `json:"target_commitish"`
		Name            interface{}   `json:"name"`
		Draft           bool          `json:"draft"`
		Prerelease      bool          `json:"prerelease"`
		CreatedAt       time.Time     `json:"created_at"`
		PublishedAt     time.Time     `json:"published_at"`
		Assets          []interface{} `json:"assets"`
		TarballUrl      string        `json:"tarball_url"`
		ZipballUrl      string        `json:"zipball_url"`
		Body            interface{}   `json:"body"`
	} `json:"release"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
}

func (r *Release) GetRelease() domain.ReleasePayload {
	return domain.ReleasePayload{
		Name: r.Release.TagName,
		Link: r.Release.HtmlUrl,
	}
}

func (r *Release) GetAction() string {
	return r.Action
}

func (r *Release) GetProject() domain.Project {
	return r.Repository.ToDomain()
}

func (r *Release) GetAuthor() domain.Author {
	return r.Sender.ToDomain()
}

func (r *Release) GetSrcBranch() domain.Branch {
	return domain.Branch{}
}

func (r *Release) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{}
}

func (r *Release) GetCommits() []domain.Commit {
	return nil
}

func (r *Release) GetCommitsAmount() int {
	return 0
}

func (r *Release) GetWorkflow() domain.Workflow {
	return domain.Workflow{}
}

func (r *Release) GetEventType() domain.EventType {
	return domain.Release
}
