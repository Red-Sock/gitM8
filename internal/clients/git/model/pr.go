package model

import (
	"time"
)

type PullRequestPayload struct {
	Action      string `json:"action"`
	Number      int    `json:"number"`
	PullRequest struct {
		Url      string `json:"url"`
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		HtmlUrl  string `json:"html_url"`
		DiffUrl  string `json:"diff_url"`
		PatchUrl string `json:"patch_url"`
		IssueUrl string `json:"issue_url"`
		Number   int    `json:"number"`
		State    string `json:"state"`
		Locked   bool   `json:"locked"`
		Title    string `json:"title"`
		User     struct {
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
		Head                Base          `json:"head"`
		Base                Base          `json:"base"`
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
	Repository struct {
		Id       int    `json:"id"`
		NodeId   string `json:"node_id"`
		Name     string `json:"name"`
		FullName string `json:"full_name"`
		Private  bool   `json:"private"`
		Owner    struct {
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
		} `json:"owner"`
		HtmlUrl                  string        `json:"html_url"`
		Description              string        `json:"description"`
		Fork                     bool          `json:"fork"`
		Url                      string        `json:"url"`
		ForksUrl                 string        `json:"forks_url"`
		KeysUrl                  string        `json:"keys_url"`
		CollaboratorsUrl         string        `json:"collaborators_url"`
		TeamsUrl                 string        `json:"teams_url"`
		HooksUrl                 string        `json:"hooks_url"`
		IssueEventsUrl           string        `json:"issue_events_url"`
		EventsUrl                string        `json:"events_url"`
		AssigneesUrl             string        `json:"assignees_url"`
		BranchesUrl              string        `json:"branches_url"`
		TagsUrl                  string        `json:"tags_url"`
		BlobsUrl                 string        `json:"blobs_url"`
		GitTagsUrl               string        `json:"git_tags_url"`
		GitRefsUrl               string        `json:"git_refs_url"`
		TreesUrl                 string        `json:"trees_url"`
		StatusesUrl              string        `json:"statuses_url"`
		LanguagesUrl             string        `json:"languages_url"`
		StargazersUrl            string        `json:"stargazers_url"`
		ContributorsUrl          string        `json:"contributors_url"`
		SubscribersUrl           string        `json:"subscribers_url"`
		SubscriptionUrl          string        `json:"subscription_url"`
		CommitsUrl               string        `json:"commits_url"`
		GitCommitsUrl            string        `json:"git_commits_url"`
		CommentsUrl              string        `json:"comments_url"`
		IssueCommentUrl          string        `json:"issue_comment_url"`
		ContentsUrl              string        `json:"contents_url"`
		CompareUrl               string        `json:"compare_url"`
		MergesUrl                string        `json:"merges_url"`
		ArchiveUrl               string        `json:"archive_url"`
		DownloadsUrl             string        `json:"downloads_url"`
		IssuesUrl                string        `json:"issues_url"`
		PullsUrl                 string        `json:"pulls_url"`
		MilestonesUrl            string        `json:"milestones_url"`
		NotificationsUrl         string        `json:"notifications_url"`
		LabelsUrl                string        `json:"labels_url"`
		ReleasesUrl              string        `json:"releases_url"`
		DeploymentsUrl           string        `json:"deployments_url"`
		CreatedAt                time.Time     `json:"created_at"`
		UpdatedAt                time.Time     `json:"updated_at"`
		PushedAt                 time.Time     `json:"pushed_at"`
		GitUrl                   string        `json:"git_url"`
		SshUrl                   string        `json:"ssh_url"`
		CloneUrl                 string        `json:"clone_url"`
		SvnUrl                   string        `json:"svn_url"`
		Homepage                 interface{}   `json:"homepage"`
		Size                     int           `json:"size"`
		StargazersCount          int           `json:"stargazers_count"`
		WatchersCount            int           `json:"watchers_count"`
		Language                 interface{}   `json:"language"`
		HasIssues                bool          `json:"has_issues"`
		HasProjects              bool          `json:"has_projects"`
		HasDownloads             bool          `json:"has_downloads"`
		HasWiki                  bool          `json:"has_wiki"`
		HasPages                 bool          `json:"has_pages"`
		HasDiscussions           bool          `json:"has_discussions"`
		ForksCount               int           `json:"forks_count"`
		MirrorUrl                interface{}   `json:"mirror_url"`
		Archived                 bool          `json:"archived"`
		Disabled                 bool          `json:"disabled"`
		OpenIssuesCount          int           `json:"open_issues_count"`
		License                  interface{}   `json:"license"`
		AllowForking             bool          `json:"allow_forking"`
		IsTemplate               bool          `json:"is_template"`
		WebCommitSignoffRequired bool          `json:"web_commit_signoff_required"`
		Topics                   []interface{} `json:"topics"`
		Visibility               string        `json:"visibility"`
		Forks                    int           `json:"forks"`
		OpenIssues               int           `json:"open_issues"`
		Watchers                 int           `json:"watchers"`
		DefaultBranch            string        `json:"default_branch"`
	} `json:"repository"`
	Organization struct {
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
	} `json:"organization"`
	Sender struct {
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
	} `json:"sender"`
}
