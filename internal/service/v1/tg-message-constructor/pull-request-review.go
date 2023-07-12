package tg_message_constructor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

func (m *MessageConstructor) extractPullRequestReview(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	if payload.GetAction() != domain.ActionSubmitted {
		return "", nil, nil
	}

	constr := constructor{}

	constr.Write(assets.Review)
	{
		author := payload.GetAuthor()
		constr.WriteWithLink(author.Name, author.Link)
	}

	constr.Write(" has reviewed pull request ")

	{
		pr := payload.GetPullRequest()

		constr.WriteWithLink(pr.Name, pr.Link)
	}

	return constr.String(), constr.format, nil
}
