package v1

import (
	"context"

	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type MessageConstructor struct {
	subscriptions dataInterfaces.Subscriptions
}

func NewMessageConstructor(repository dataInterfaces.Repository) *MessageConstructor {
	return &MessageConstructor{
		subscriptions: repository.Subscriptions(),
	}
}

func (m *MessageConstructor) Parse(in domain.TicketRequest) ([]interfaces.MessageOut, error) {
	msg, err := m.extractMessage(in.Payload)
	if err != nil {
		return nil, errors.Wrap(err, "error creating message for event")
	}

	subs, err := m.subscriptions.GetSubscribers(context.Background(), in.TicketId)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining subscribers from repo")
	}

	out := make([]interfaces.MessageOut, 0, len(subs))

	for _, sub := range subs {
		out = append(out, &response.MessageOut{
			ChatId: int64(sub.ChatId),
			Text:   msg,
		})
	}

	return out, nil
}

func (m *MessageConstructor) extractMessage(payload domain.Payload) (string, error) {
	switch payload.GetEventType() {
	case domain.Push:
		return payload.GetAuthor().Name + " has pushed to " + payload.GetProject().Name + ".\n" +
			"Branch " + payload.GetSrcBranch().Name, nil
	case domain.Ping:
		// TODO GITM-8
		return payload.GetProject().Name + " successfully has been pinged!", nil
	case domain.Comment:
		// TODO GITM-7
		return payload.GetAuthor().Name + " left a comment at " + payload.GetPullRequest().Name, nil
	case domain.PullRequest:
		// TODO GITM-9
		return payload.GetAuthor().Name + " created a pull request", nil
	case domain.Release:
		// TODO GITM-13
		return "New version of " + payload.GetProject().Name + " has been released", nil
	case domain.WorkflowJob:
		// TODO GITM-15
		return "Event messaging is still in progress. Await GITM-15 to be done", nil
	case domain.WorkflowRun:
		// TODO GITM-16
		return "Event messaging is still in progress. Await GITM-16 to be done", nil
	case domain.WorkflowManualStart:
		// TODO GITM-17
		return "Event messaging is still in progress. Await GITM-17 to be done", nil
	default:
		return "", domain.ErrUnknownEventType
	}
}
