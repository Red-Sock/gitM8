package domain

type Request struct {
	Src     RepoType
	Type    EventType
	Payload []byte
}

type EventType int

const (
	Invalid EventType = iota

	Comment
	Ping

	PullRequest
	PullRequestComment
	PullRequestReview

	Push
	Release

	CommitStatus
	WorkflowJob
	WorkflowRun
	WorkflowManualStart
)

func (w *EventType) ParseGithub(in string) {
	*w, _ = githubEventsToDomain[in]
}

func (w *EventType) String() string {
	switch *w {
	case Comment:
		return "Commentary"
	case Ping:
		return "Webhook ping"
	case PullRequest:
		return "Pull request"
	case PullRequestComment:
		return "Comment on pull request"
	case PullRequestReview:
		return "Review on pull request"
	case Push:
		return "Push to repository"
	case Release:
		return "New release"
	case CommitStatus:
		return "Commit status"
	case WorkflowJob:
		return "Workflow job info"
	case WorkflowRun:
		return "Workflow info"
	case WorkflowManualStart:
		return "Workflow manual start"
	default:
		return "invalid type"
	}
}

var githubEventsToDomain = map[string]EventType{
	"commit_comment":              Comment,
	"ping":                        Ping,
	"pull_request":                PullRequest,
	"pull_request_review_comment": PullRequestComment,
	"pull_request_review":         PullRequestReview,
	"push":                        Push,
	"release":                     Release,
	"status":                      CommitStatus,
	"workflow_job":                WorkflowJob,
	"workflow_run":                WorkflowRun,
	"workflow_dispatch":           WorkflowManualStart,
}
