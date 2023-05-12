package model

import (
	"time"
)

// Assignee contains GitHub's assignee information
type Assignee struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Milestone contains GitHub's milestone information
type Milestone struct {
	URL         string `json:"url"`
	HTMLURL     string `json:"html_url"`
	LabelsURL   string `json:"labels_url"`
	ID          int64  `json:"id"`
	NodeID      string `json:"node_id"`
	Number      int64  `json:"number"`
	State       string `json:"state"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Creator     struct {
		Login             string `json:"login"`
		ID                int64  `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"creator"`
	OpenIssues   int64     `json:"open_issues"`
	ClosedIssues int64     `json:"closed_issues"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	ClosedAt     time.Time `json:"closed_at"`
	DueOn        time.Time `json:"due_on"`
}

// MergedBy contains GitHub's merged-by information
type MergedBy struct {
	Login             string `json:"login"`
	ID                int64  `json:"id"`
	NodeID            string `json:"node_id"`
	AvatarURL         string `json:"avatar_url"`
	GravatarID        string `json:"gravatar_id"`
	URL               string `json:"url"`
	HTMLURL           string `json:"html_url"`
	FollowersURL      string `json:"followers_url"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	OrganizationsURL  string `json:"organizations_url"`
	ReposURL          string `json:"repos_url"`
	EventsURL         string `json:"events_url"`
	ReceivedEventsURL string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
}

// Asset contains GitHub's asset information
type Asset struct {
	URL                string    `json:"url"`
	BrowserDownloadURL string    `json:"browser_download_url"`
	ID                 int64     `json:"id"`
	NodeID             string    `json:"node_id"`
	Name               string    `json:"name"`
	Label              string    `json:"label"`
	State              string    `json:"state"`
	ContentType        string    `json:"content_type"`
	Size               int64     `json:"size"`
	DownloadCount      int64     `json:"download_count"`
	CreatedAt          time.Time `json:"created_at"`
	UpdatedAt          time.Time `json:"updated_at"`
	Uploader           struct {
		Login             string `json:"login"`
		ID                int64  `json:"id"`
		NodeID            string `json:"node_id"`
		AvatarURL         string `json:"avatar_url"`
		GravatarID        string `json:"gravatar_id"`
		URL               string `json:"url"`
		HTMLURL           string `json:"html_url"`
		FollowersURL      string `json:"followers_url"`
		FollowingURL      string `json:"following_url"`
		GistsURL          string `json:"gists_url"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		OrganizationsURL  string `json:"organizations_url"`
		ReposURL          string `json:"repos_url"`
		EventsURL         string `json:"events_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		Type              string `json:"type"`
		SiteAdmin         bool   `json:"site_admin"`
	} `json:"uploader"`
}

// Parent contains GitHub's parent information
type Parent struct {
	URL string `json:"url"`
	Sha string `json:"sha"`
}

// Label contains Issue's Label information
type Label struct {
	ID      int64  `json:"id"`
	NodeID  string `json:"node_id"`
	URL     string `json:"url"`
	Name    string `json:"name"`
	Color   string `json:"color"`
	Default bool   `json:"default"`
}

// Team contains GitHub's Team information
type Team struct {
	Name            string `json:"name"`
	ID              int64  `json:"id"`
	NodeID          string `json:"node_id"`
	Slug            string `json:"slug"`
	Permission      string `json:"permission"`
	URL             string `json:"url"`
	MembersURL      string `json:"members_url"`
	RepositoriesURL string `json:"repositories_url"`
	Parent          *Team  `json:"parent,omitempty"`
}

// Step contains workflow_job step information
type Step struct {
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Conclusion  string    `json:"conclusion"`
	ID          int64     `json:"id"`
	StartedAt   time.Time `json:"started_at"`
	CompletedAt time.Time `json:"completed_at"`
}
