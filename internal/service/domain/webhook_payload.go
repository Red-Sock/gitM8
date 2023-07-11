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
	ActionUnknown   = "unknown"
	ActionSubmitted = "submitted"
	ActionCreated   = "created"
)
