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
		// TODO GITM-8
	case domain.Comment:
		// TODO GITM-7
	case domain.PullRequest:
		// TODO GITM-9
	case domain.Release:
		// TODO GITM-13
	case domain.WorkflowJob:
		// TODO GITM-15
	case domain.WorkflowRun:
		// TODO GITM-16
	case domain.WorkflowManualStart:
		// TODO GITM-17
	default:
		return nil, domain.ErrUnknownEventType
	}

	err := json.Unmarshal(payload, out)
	if err != nil {
		return nil, errors.Wrap(err, "error unmarshalling payload to structure")
	}
	return out, nil
}
