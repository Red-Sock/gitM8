package domain

type Request struct {
	Src     RepoType
	Type    Type
	Payload []byte
}

type Type int

const (
	Invalid Type = iota

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
	WorkflowStart
)

func (w *Type) ParseGithub(in string) {
	*w, _ = githubEventsToDomain[in]
}

var githubEventsToDomain = map[string]Type{
	"push":                        Push,
	"pull_request":                PullRequest,
	"commit_comment":              Comment,
	"ping":                        Ping,
	"pull_request_review_comment": PullRequestComment,
	"pull_request_review":         PullRequestReview,
	"release":                     Release,
	"status":                      CommitStatus,
	"workflow_job":                WorkflowJob,
	"workflow_run":                WorkflowRun,
	"workflow_dispatch":           WorkflowStart,
}
