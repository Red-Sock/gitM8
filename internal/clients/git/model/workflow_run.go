package model

import (
	"time"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

// WorkflowPayload - pipeline containing jobs
type WorkflowPayload struct {
	Action      string `json:"action"`
	WorkflowRun struct {
		Id               int64         `json:"id"`
		Name             string        `json:"name"`
		NodeId           string        `json:"node_id"`
		HeadBranch       string        `json:"head_branch"`
		HeadSha          string        `json:"head_sha"`
		Path             string        `json:"path"`
		DisplayTitle     string        `json:"display_title"`
		RunNumber        int           `json:"run_number"`
		Event            string        `json:"event"`
		Status           string        `json:"status"`
		Conclusion       string        `json:"conclusion"`
		WorkflowId       int           `json:"workflow_id"`
		CheckSuiteId     int64         `json:"check_suite_id"`
		CheckSuiteNodeId string        `json:"check_suite_node_id"`
		Url              string        `json:"url"`
		HtmlUrl          string        `json:"html_url"`
		PullRequests     []interface{} `json:"pull_requests"`
		CreatedAt        time.Time     `json:"created_at"`
		UpdatedAt        time.Time     `json:"updated_at"`
		Actor            struct {
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
		} `json:"actor"`
		RunAttempt          int           `json:"run_attempt"`
		ReferencedWorkflows []interface{} `json:"referenced_workflows"`
		RunStartedAt        time.Time     `json:"run_started_at"`
		TriggeringActor     struct {
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
		} `json:"triggering_actor"`
		JobsUrl            string      `json:"jobs_url"`
		LogsUrl            string      `json:"logs_url"`
		CheckSuiteUrl      string      `json:"check_suite_url"`
		ArtifactsUrl       string      `json:"artifacts_url"`
		CancelUrl          string      `json:"cancel_url"`
		RerunUrl           string      `json:"rerun_url"`
		PreviousAttemptUrl interface{} `json:"previous_attempt_url"`
		WorkflowUrl        string      `json:"workflow_url"`
		HeadCommit         struct {
			Id        string    `json:"id"`
			TreeId    string    `json:"tree_id"`
			Message   string    `json:"message"`
			Timestamp time.Time `json:"timestamp"`
			Author    struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"author"`
			Committer struct {
				Name  string `json:"name"`
				Email string `json:"email"`
			} `json:"committer"`
		} `json:"head_commit"`
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
			HtmlUrl          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			Url              string `json:"url"`
			ForksUrl         string `json:"forks_url"`
			KeysUrl          string `json:"keys_url"`
			CollaboratorsUrl string `json:"collaborators_url"`
			TeamsUrl         string `json:"teams_url"`
			HooksUrl         string `json:"hooks_url"`
			IssueEventsUrl   string `json:"issue_events_url"`
			EventsUrl        string `json:"events_url"`
			AssigneesUrl     string `json:"assignees_url"`
			BranchesUrl      string `json:"branches_url"`
			TagsUrl          string `json:"tags_url"`
			BlobsUrl         string `json:"blobs_url"`
			GitTagsUrl       string `json:"git_tags_url"`
			GitRefsUrl       string `json:"git_refs_url"`
			TreesUrl         string `json:"trees_url"`
			StatusesUrl      string `json:"statuses_url"`
			LanguagesUrl     string `json:"languages_url"`
			StargazersUrl    string `json:"stargazers_url"`
			ContributorsUrl  string `json:"contributors_url"`
			SubscribersUrl   string `json:"subscribers_url"`
			SubscriptionUrl  string `json:"subscription_url"`
			CommitsUrl       string `json:"commits_url"`
			GitCommitsUrl    string `json:"git_commits_url"`
			CommentsUrl      string `json:"comments_url"`
			IssueCommentUrl  string `json:"issue_comment_url"`
			ContentsUrl      string `json:"contents_url"`
			CompareUrl       string `json:"compare_url"`
			MergesUrl        string `json:"merges_url"`
			ArchiveUrl       string `json:"archive_url"`
			DownloadsUrl     string `json:"downloads_url"`
			IssuesUrl        string `json:"issues_url"`
			PullsUrl         string `json:"pulls_url"`
			MilestonesUrl    string `json:"milestones_url"`
			NotificationsUrl string `json:"notifications_url"`
			LabelsUrl        string `json:"labels_url"`
			ReleasesUrl      string `json:"releases_url"`
			DeploymentsUrl   string `json:"deployments_url"`
		} `json:"repository"`
		HeadRepository struct {
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
			HtmlUrl          string `json:"html_url"`
			Description      string `json:"description"`
			Fork             bool   `json:"fork"`
			Url              string `json:"url"`
			ForksUrl         string `json:"forks_url"`
			KeysUrl          string `json:"keys_url"`
			CollaboratorsUrl string `json:"collaborators_url"`
			TeamsUrl         string `json:"teams_url"`
			HooksUrl         string `json:"hooks_url"`
			IssueEventsUrl   string `json:"issue_events_url"`
			EventsUrl        string `json:"events_url"`
			AssigneesUrl     string `json:"assignees_url"`
			BranchesUrl      string `json:"branches_url"`
			TagsUrl          string `json:"tags_url"`
			BlobsUrl         string `json:"blobs_url"`
			GitTagsUrl       string `json:"git_tags_url"`
			GitRefsUrl       string `json:"git_refs_url"`
			TreesUrl         string `json:"trees_url"`
			StatusesUrl      string `json:"statuses_url"`
			LanguagesUrl     string `json:"languages_url"`
			StargazersUrl    string `json:"stargazers_url"`
			ContributorsUrl  string `json:"contributors_url"`
			SubscribersUrl   string `json:"subscribers_url"`
			SubscriptionUrl  string `json:"subscription_url"`
			CommitsUrl       string `json:"commits_url"`
			GitCommitsUrl    string `json:"git_commits_url"`
			CommentsUrl      string `json:"comments_url"`
			IssueCommentUrl  string `json:"issue_comment_url"`
			ContentsUrl      string `json:"contents_url"`
			CompareUrl       string `json:"compare_url"`
			MergesUrl        string `json:"merges_url"`
			ArchiveUrl       string `json:"archive_url"`
			DownloadsUrl     string `json:"downloads_url"`
			IssuesUrl        string `json:"issues_url"`
			PullsUrl         string `json:"pulls_url"`
			MilestonesUrl    string `json:"milestones_url"`
			NotificationsUrl string `json:"notifications_url"`
			LabelsUrl        string `json:"labels_url"`
			ReleasesUrl      string `json:"releases_url"`
			DeploymentsUrl   string `json:"deployments_url"`
		} `json:"head_repository"`
	} `json:"workflow_run"`
	Workflow struct {
		Id        int       `json:"id"`
		NodeId    string    `json:"node_id"`
		Name      string    `json:"name"`
		Path      string    `json:"path"`
		State     string    `json:"state"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
		Url       string    `json:"url"`
		HtmlUrl   string    `json:"html_url"`
		BadgeUrl  string    `json:"badge_url"`
	} `json:"workflow"`
	Repository   Repo         `json:"repository"`
	Organization Organization `json:"organization"`
	Sender       User         `json:"sender"`
}

func (w *WorkflowPayload) GetWorkflow() domain.Workflow {
	return domain.Workflow{
		Name:   w.WorkflowRun.Name,
		Link:   w.WorkflowRun.HtmlUrl,
		Status: domain.WorkflowStatus(w.WorkflowRun.Status),
		Result: domain.WorkflowResult(w.WorkflowRun.Conclusion),
	}
}

func (w *WorkflowPayload) GetAction() string {
	return w.Action
}

func (w *WorkflowPayload) GetProject() domain.Project {
	return w.Repository.ToDomain()
}

func (w *WorkflowPayload) GetAuthor() domain.Author {
	return w.Sender.ToDomain()
}

func (w *WorkflowPayload) GetSrcBranch() domain.Branch {
	return domain.Branch{}
}

func (w *WorkflowPayload) GetPullRequest() domain.PullRequestPayload {
	return domain.PullRequestPayload{}
}

func (w *WorkflowPayload) GetCommits() []domain.Commit {
	return nil
}

func (w *WorkflowPayload) GetCommitsAmount() int {
	return 0
}

func (w *WorkflowPayload) GetEventType() domain.EventType {
	return domain.WorkflowRun
}

func (w *WorkflowPayload) GetRelease() domain.ReleasePayload {
	return domain.ReleasePayload{}
}
