package domain

type Payload interface {
	GetEventType() EventType

	GetAction() Action
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

type Action string

const (
	ActionUnknown      Action = "unknown"
	ActionSubmitted    Action = "submitted"
	ActionCreated      Action = "created"
	ActionOpened       Action = "opened"
	ActionSynchronized Action = "synchronize"
)

type WorkflowStatus string

const (
	WorkflowStatusInProgress WorkflowStatus = "in_progress"
	WorkflowStatusCompleted  WorkflowStatus = "completed"
)

type WorkflowResult string

const (
	WorkflowResultSuccess = "success"
	WorkflowResultFailure = "failure"
)
