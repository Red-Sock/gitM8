package model

import (
	"encoding/json"

	"github.com/pkg/errors"
	"github.com/tidwall/gjson"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

func SelectModel(eventType domain.EventType, payload []byte) (domain.Payload, error) {
	var out domain.Payload

	switch eventType {
	case domain.Push:
		out = &PushPayload{}
	case domain.Ping:
		out = &PingPayload{}
	case domain.Comment:
		parsed := gjson.ParseBytes(payload)
		if parsed.Get("issue").Exists() {
			out = &PullRequestComment{}
		} else {
			out = &PullRequestCodeComment{}
		}
	case domain.PullRequest:
		out = &PullRequestPayload{}
	case domain.Release:
		// TODO GITM-13
		return nil, domain.ErrUnknownEventType
	case domain.WorkflowJob:
		// TODO GITM-15
		return nil, domain.ErrUnknownEventType
	case domain.WorkflowRun:
		// TODO GITM-16
		return nil, domain.ErrUnknownEventType
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
