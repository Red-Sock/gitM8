package v1

import (
	"context"
	"strconv"
	"strings"

	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/response"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

type constructors func(payload domain.Payload) (string, []tgbotapi.MessageEntity, error)

type MessageConstructor struct {
	subscriptions           dataInterfaces.Subscriptions
	eventTypeToConstructors map[domain.EventType]constructors
}

func NewMessageConstructor(repository dataInterfaces.Repository) *MessageConstructor {
	m := &MessageConstructor{
		subscriptions: repository.Subscriptions(),
	}

	m.eventTypeToConstructors = map[domain.EventType]constructors{
		domain.Push: m.extractPushMessage,
		domain.Ping: m.extractPingMessage,
	}

	return m
}

func (m *MessageConstructor) Parse(in domain.TicketRequest) ([]interfaces.MessageOut, error) {
	construct, ok := m.eventTypeToConstructors[in.Payload.GetEventType()]
	if !ok {
		return nil, domain.ErrUnknownEventType
	}

	subs, err := m.subscriptions.GetSubscribers(context.Background(), in.TicketId)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining subscribers from repo")
	}

	msg, format, err := construct(in.Payload)
	if err != nil {
		return nil, err
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

func (m *MessageConstructor) extractPushMessage(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	constr.Write(assets.Push)
	{
		author := payload.GetAuthor()
		constr.WriteWithLink(author.Name, author.Link)
	}
	constr.Write(" has pushed to ")

	{
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}

	{
		srcBranch := payload.GetSrcBranch()
		constr.Write(" to branch ")
		constr.WriteWithLink(srcBranch.Name, srcBranch.Link)
	}

	{
		commits := payload.GetCommits()
		constr.Write(" " + strconv.Itoa(len(commits)))
		if len(commits) == 1 {
			constr.Write(" commit")
		} else {
			constr.Write(" commits")
		}

	}

	return constr.String(), constr.format, nil
}

func (m *MessageConstructor) extractPingMessage(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	constr.Write("Repository ")
	{
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}
	constr.Write(" has pinged this webhook!\n")
	constr.Write("Sending a pong right away")
	// for some reason assets.Ping causes constructor to mess with index of format
	// for example, assets.Ping putted before link, causes link to shift one symbol left
	constr.Write(assets.Ping)

	return constr.String(), constr.format, nil
}

type constructor struct {
	text      []string
	format    []tgbotapi.MessageEntity
	idx       int
	separator string
}

func (c *constructor) Write(text string) {
	c.idx += len([]rune(text))
	c.text = append(c.text, text)
}

func (c *constructor) WriteWithFormat(text, format string) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   format,
		Offset: c.idx,
		Length: len(text),
	})

	c.Write(text)
}

func (c *constructor) WriteWithLink(text, url string) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   response.TextLinkTextFormat,
		Offset: c.idx,
		Length: len(text),
		URL:    url,
	})

	c.Write(text)
}

func (c *constructor) WriteWithMention(text string, userID uint64) {
	c.format = append(c.format, tgbotapi.MessageEntity{
		Type:   response.MentionTextFormat,
		Offset: c.idx,
		Length: len(text),
		User: &tgbotapi.User{
			ID: int64(userID),
		},
	})

	c.Write(text)
}

func (c *constructor) String() string {
	return strings.Join(c.text, c.separator)
}
