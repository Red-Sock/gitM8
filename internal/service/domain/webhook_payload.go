package domain

type Payload interface {
	GetEventType() EventType

	GetAction() string
	// GetProject - obtain project in which event has happened
	GetProject() Project
	// GetAuthor - obtain user who caused action
	GetAuthor() Author
	// GetSrcBranch - source branch from which caused event
	GetSrcBranch() Branch
	// GetPullRequest - obtain pull request information (PullRequest)
	GetPullRequest() PullRequestPayload
	// GetCommits - returns commits (Push)
	GetCommits() []Commit
	// GetCommitsAmount - returns amount of commits (PullRequest)
	GetCommitsAmount() int
	// GetWorkflow - returns pipeline data (WorkflowRun)
	GetWorkflow() Workflow
	// GetRelease - returns release information (Release)
	GetRelease() ReleasePayload
}

type Project struct {
	Name string
	Link string
}

type Author struct {
	Name string
	Link string
}

type Branch struct {
	Name string
	Link string
}

type Workflow struct {
	Name   string
	Link   string
	Status WorkflowStatus
	Result WorkflowResult
}

type ReleasePayload struct {
	Name string
	Link string
}

type PullRequestState int

const (
	PullRequestStateUnknown PullRequestState = iota
	PullRequestStateOpened
	PullRequestStateClosed
)

type PullRequestPayload struct {
	StateStr string
	Name     string
	Link     string
	Base     Branch
	Target   Branch
}

func (p *PullRequestPayload) GetState() PullRequestState {
	switch p.StateStr {
	case "open":
		return PullRequestStateOpened
	case "closed":
		return PullRequestStateClosed

	default:
		return PullRequestStateUnknown
	}
}

type Commit struct {
	Author Author
}

const (
	ActionUnknown      = "unknown"
	ActionSubmitted    = "submitted"
	ActionCreated      = "created"
	ActionOpened       = "opened"
	ActionSynchronized = "synchronize"

	Deleted     = "deleted"
	Edited      = "edited"
	PreReleased = "prereleased"
	Published   = "published"
	Released    = "released"
	UnPublished = "unpublished"
)

type WorkflowStatus string

const (
	WorkflowStatusInProgress WorkflowStatus = "in_progress"
	WorkflowStatusCompleted  WorkflowStatus = "completed"
)

type WorkflowResult string

const (
	WorkflowResultSuccess WorkflowResult = "success"
	WorkflowResultFailure WorkflowResult = "failure"
)
