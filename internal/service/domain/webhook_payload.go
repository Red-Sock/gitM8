package domain

type Payload interface {
	GetEventType() EventType

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
)

type PullRequestPayload struct {
	StateStr string
	Name     string
	Link     string
}

func (p *PullRequestPayload) GetState() PullRequestState {
	switch p.StateStr {
	case "open":
		return PullRequestStateOpened
	default:
		return PullRequestStateUnknown
	}
}

type Commit struct {
	Author Author
}
