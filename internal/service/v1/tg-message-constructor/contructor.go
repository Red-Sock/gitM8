package tg_message_constructor

import (
	"context"

	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/response"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type constructors func(payload domain.Payload) (string, []tgbotapi.MessageEntity, error)

type MessageConstructor struct {
	subscriptions           dataInterfaces.Subscriptions
	eventTypeToConstructors map[domain.EventType]constructors

	workflowConstructor
	releaseConstructor
}

func NewMessageConstructor(repository dataInterfaces.Repository) *MessageConstructor {
	m := &MessageConstructor{
		subscriptions: repository.Subscriptions(),
	}

	m.eventTypeToConstructors = map[domain.EventType]constructors{
		domain.Ping:          m.extractPingMessage,
		domain.Push:          m.extractPushMessage,
		domain.PullRequest:   m.extractPullRequest,
		domain.ReviewComment: m.extractPullRequestReview,
		domain.IssueComment:  m.extractIssueComment,
		domain.WorkflowRun:   m.extractWorkflowRun,
		domain.Release:       m.extractRelease,
	}

	return m
}

func (m *MessageConstructor) Parse(ctx context.Context, in domain.TicketRequest) ([]interfaces.MessageOut, error) {
	construct, ok := m.eventTypeToConstructors[in.Payload.GetEventType()]
	if !ok {
		return nil, domain.ErrUnknownEventType
	}

	subs, err := m.subscriptions.GetSubscribers(ctx, in.TicketId)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining subscribers from repo")
	}

	msg, format, err := construct(in.Payload)
	if err != nil {
		return nil, err
	}

	if msg == "" {
		return nil, nil
	}

	out := make([]interfaces.MessageOut, 0, len(subs))

	for _, sub := range subs {
		out = append(out, &response.MessageOut{
			ChatId:   int64(sub.ChatId),
			Text:     msg,
			Entities: format,
		})
	}

	return out, nil
}
