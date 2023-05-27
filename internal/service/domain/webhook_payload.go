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
}

type Project struct {
	Name string
}

type Author struct {
	Name string
}

type Branch struct {
	Name string
}

type PullRequestPayload struct {
	Name string
}
