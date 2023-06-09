package model

import (
	"encoding/json"

	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

func SelectModel(eventType domain.EventType, payload []byte) (domain.Payload, error) {
	var out domain.Payload

	switch eventType {
	case domain.Push:
		out = &PushPayload{}
	case domain.Ping:
		out = &PingPayload{}
	case domain.ReviewComment:
		out = &PullRequestCodeComment{}
	case domain.IssueComment:
		out = &PullRequestComment{}
	case domain.PullRequest:
		out = &PullRequestPayload{}
	case domain.Release:
		// TODO GITM-13
		return nil, domain.ErrUnknownEventType
	case domain.WorkflowJob:
		// TODO GITM-15
		return nil, domain.ErrUnknownEventType
	case domain.WorkflowRun:
		out = &WorkflowPayload{}
	case domain.WorkflowManualStart:
		// TODO GITM-17
		return nil, domain.ErrUnknownEventType
	default:
		return nil, domain.ErrUnknownEventType
	}

	err := json.Unmarshal(payload, out)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling payload to structure")
	}

	return out, nil
}
