package domain

import (
	"github.com/pkg/errors"
)

type EventType int

var ErrUnknownEventType = errors.New("unknown event type")

const (
	Invalid EventType = iota
	// Push - push to branch
	Push
	// Ping - healthcheck webhook
	Ping
	// Comment - comment in pr(not connected to code, connected to code and review (summary of comments to code))
	Comment
	// PullRequest - pr connected action
	PullRequest

	// Release - action connected to realising new code version (e.g. tag creation)
	Release

	WorkflowJob
	WorkflowRun
	WorkflowManualStart
)

func (w *EventType) ParseGithub(in string) {
	*w, _ = githubEventsToDomain[in]
}

func (w *EventType) String() string {
	switch *w {
	case Ping:
		return "Webhook healthcheck"
	case Comment:
		return "Pull request comments"
	case PullRequest:
		return "Pull request actions"
	case Push:
		return "Pushes to branch"
	case Release:
		return "Code releases"
	case WorkflowJob:
		return "Workflow job info"
	case WorkflowRun:
		return "Workflow info"
	case WorkflowManualStart:
		return "Workflow manual start notifications"
	default:
		return "invalid type"
	}
}

var githubEventsToDomain = map[string]EventType{
	"ping": Ping,
	"push": Push,

	"pull_request":                PullRequest,
	"issue_comment":               Comment,
	"pull_request_review_comment": Comment,
	"pull_request_review":         Comment,

	"release": Release,

	"workflow_job":      WorkflowJob,
	"workflow_run":      WorkflowRun,
	"workflow_dispatch": WorkflowManualStart,
}

func GetEventTypes() []EventType {
	return []EventType{
		Push,
		Ping,
		Comment,
		PullRequest,
		Release,
		WorkflowJob,
		WorkflowRun,
		WorkflowManualStart,
	}
}
