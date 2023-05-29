package model

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type Base struct {
	Label string `json:"label"`
	Ref   string `json:"ref"`
	Sha   string `json:"sha"`
	User  User   `json:"user"`
	Repo  Repo
}

type User struct {
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
}

func (u *User) ToDomain() domain.Author {
	return domain.Author{
		Name: u.Login,
		Link: u.HtmlUrl,
	}
}

type UserShort struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

func (u *UserShort) ToDomain() domain.Author {
	return domain.Author{
		Name: u.Username,
		Link: "https://github.com/" + u.Username,
	}
}

type Organization struct {
	Login            string `json:"login"`
	Id               int    `json:"id"`
	NodeId           string `json:"node_id"`
	Url              string `json:"url"`
	ReposUrl         string `json:"repos_url"`
	EventsUrl        string `json:"events_url"`
	HooksUrl         string `json:"hooks_url"`
	IssuesUrl        string `json:"issues_url"`
	MembersUrl       string `json:"members_url"`
	PublicMembersUrl string `json:"public_members_url"`
	AvatarUrl        string `json:"avatar_url"`
	Description      string `json:"description"`
}

type Repo struct {
	Id                        int           `json:"id"`
	NodeId                    string        `json:"node_id"`
	Name                      string        `json:"name"`
	FullName                  string        `json:"full_name"`
	Private                   bool          `json:"private"`
	Owner                     User          `json:"owner"`
	HtmlUrl                   string        `json:"html_url"`
	Description               string        `json:"description"`
	Fork                      bool          `json:"fork"`
	Url                       string        `json:"url"`
	ForksUrl                  string        `json:"forks_url"`
	KeysUrl                   string        `json:"keys_url"`
	CollaboratorsUrl          string        `json:"collaborators_url"`
	TeamsUrl                  string        `json:"teams_url"`
	HooksUrl                  string        `json:"hooks_url"`
	IssueEventsUrl            string        `json:"issue_events_url"`
	EventsUrl                 string        `json:"events_url"`
	AssigneesUrl              string        `json:"assignees_url"`
	BranchesUrl               string        `json:"branches_url"`
	TagsUrl                   string        `json:"tags_url"`
	BlobsUrl                  string        `json:"blobs_url"`
	GitTagsUrl                string        `json:"git_tags_url"`
	GitRefsUrl                string        `json:"git_refs_url"`
	TreesUrl                  string        `json:"trees_url"`
	StatusesUrl               string        `json:"statuses_url"`
	LanguagesUrl              string        `json:"languages_url"`
	StargazersUrl             string        `json:"stargazers_url"`
	ContributorsUrl           string        `json:"contributors_url"`
	SubscribersUrl            string        `json:"subscribers_url"`
	SubscriptionUrl           string        `json:"subscription_url"`
	CommitsUrl                string        `json:"commits_url"`
	GitCommitsUrl             string        `json:"git_commits_url"`
	CommentsUrl               string        `json:"comments_url"`
	IssueCommentUrl           string        `json:"issue_comment_url"`
	ContentsUrl               string        `json:"contents_url"`
	CompareUrl                string        `json:"compare_url"`
	MergesUrl                 string        `json:"merges_url"`
	ArchiveUrl                string        `json:"archive_url"`
	DownloadsUrl              string        `json:"downloads_url"`
	IssuesUrl                 string        `json:"issues_url"`
	PullsUrl                  string        `json:"pulls_url"`
	MilestonesUrl             string        `json:"milestones_url"`
	NotificationsUrl          string        `json:"notifications_url"`
	LabelsUrl                 string        `json:"labels_url"`
	ReleasesUrl               string        `json:"releases_url"`
	DeploymentsUrl            string        `json:"deployments_url"`
	CreatedAt                 Time          `json:"created_at"`
	UpdatedAt                 Time          `json:"updated_at"`
	PushedAt                  Time          `json:"pushed_at"`
	GitUrl                    string        `json:"git_url"`
	SshUrl                    string        `json:"ssh_url"`
	CloneUrl                  string        `json:"clone_url"`
	SvnUrl                    string        `json:"svn_url"`
	Homepage                  interface{}   `json:"homepage"`
	Size                      int           `json:"size"`
	StargazersCount           int           `json:"stargazers_count"`
	WatchersCount             int           `json:"watchers_count"`
	Language                  interface{}   `json:"language"`
	HasIssues                 bool          `json:"has_issues"`
	HasProjects               bool          `json:"has_projects"`
	HasDownloads              bool          `json:"has_downloads"`
	HasWiki                   bool          `json:"has_wiki"`
	HasPages                  bool          `json:"has_pages"`
	HasDiscussions            bool          `json:"has_discussions"`
	ForksCount                int           `json:"forks_count"`
	MirrorUrl                 interface{}   `json:"mirror_url"`
	Archived                  bool          `json:"archived"`
	Disabled                  bool          `json:"disabled"`
	OpenIssuesCount           int           `json:"open_issues_count"`
	License                   interface{}   `json:"license"`
	AllowForking              bool          `json:"allow_forking"`
	IsTemplate                bool          `json:"is_template"`
	WebCommitSignoffRequired  bool          `json:"web_commit_signoff_required"`
	Topics                    []interface{} `json:"topics"`
	Visibility                string        `json:"visibility"`
	Forks                     int           `json:"forks"`
	OpenIssues                int           `json:"open_issues"`
	Watchers                  int           `json:"watchers"`
	DefaultBranch             string        `json:"default_branch"`
	AllowSquashMerge          bool          `json:"allow_squash_merge"`
	AllowMergeCommit          bool          `json:"allow_merge_commit"`
	AllowRebaseMerge          bool          `json:"allow_rebase_merge"`
	AllowAutoMerge            bool          `json:"allow_auto_merge"`
	DeleteBranchOnMerge       bool          `json:"delete_branch_on_merge"`
	AllowUpdateBranch         bool          `json:"allow_update_branch"`
	UseSquashPrTitleAsDefault bool          `json:"use_squash_pr_title_as_default"`
	SquashMergeCommitMessage  string        `json:"squash_merge_commit_message"`
	SquashMergeCommitTitle    string        `json:"squash_merge_commit_title"`
	MergeCommitMessage        string        `json:"merge_commit_message"`
	MergeCommitTitle          string        `json:"merge_commit_title"`
}

type Link struct {
	Href string `json:"href"`
}

type Links struct {
	Self           Link `json:"self"`
	Html           Link `json:"html"`
	Issue          Link `json:"issue"`
	Comments       Link `json:"comments"`
	ReviewComments Link `json:"review_comments"`
	ReviewComment  Link `json:"review_comment"`
	Commits        Link `json:"commits"`
	Statuses       Link `json:"statuses"`
	PullRequest    Link `json:"pull_request"`
}

type Commit struct {
	Id        string        `json:"id"`
	TreeId    string        `json:"tree_id"`
	Distinct  bool          `json:"distinct"`
	Message   string        `json:"message"`
	Timestamp Time          `json:"timestamp"`
	Url       string        `json:"url"`
	Author    UserShort     `json:"author"`
	Committer UserShort     `json:"committer"`
	Added     []interface{} `json:"added"`
	Removed   []interface{} `json:"removed"`
	Modified  []string      `json:"modified"`
}

type Time time.Time

func (t *Time) UnmarshalJSON(in []byte) error {
	i, err := strconv.ParseInt(string(in), 10, 64)
	if err == nil {
		*t = Time(time.Unix(i, 0))
		return nil
	}
	// 2006-01-02T15:04:05Z07:00
	var newT time.Time
	err = json.Unmarshal(in, &newT)
	if err == nil {
		*t = Time(newT)
	}

	return err
}
