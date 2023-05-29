package domain

type Payload interface {
	GetEventType() EventType

	// GetProject - obtain project in which event has happened
	GetProject() Project
	// GetAuthor - obtain user who caused action
	GetAuthor() Author
	// GetSrcBranch - source branch from which caused event
	GetSrcBranch() Branch
	// GetPullRequest - obtain pull request information
	GetPullRequest() PullRequestPayload
	// GetCommits - returns commits
	GetCommits() []Commit
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

type PullRequestPayload struct {
	Name string
}

type Commit struct {
	Author Author
}
