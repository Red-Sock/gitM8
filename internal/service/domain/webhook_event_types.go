package domain

import (
	"github.com/pkg/errors"
)

type EventType int

var (
	ErrUnknownEventType      = errors.New("unknown event type")
	ErrInvalidActionForEvent = errors.New("invalid action for event")
)

const (
	Invalid EventType = iota
	// Push - push to branch
	Push
	// Ping - healthcheck webhook
	Ping
	// ReviewComment - comment in pr(not connected to code, connected to code and review (summary of comments to code))
	ReviewComment
	IssueComment

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
	case ReviewComment:
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

	"pull_request":        PullRequest,
	"pull_request_review": ReviewComment,
	"issue_comment":       IssueComment,

	"release": Release,

	"workflow_job":      WorkflowJob,
	"workflow_run":      WorkflowRun,
	"workflow_dispatch": WorkflowManualStart,
}

func GetEventTypes() []EventType {
	return []EventType{
		Push,
		Ping,
		ReviewComment,
		PullRequest,
		Release,
		WorkflowJob,
		WorkflowRun,
		WorkflowManualStart,
	}
}
