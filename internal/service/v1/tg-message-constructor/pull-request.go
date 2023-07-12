package tg_message_constructor

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

func (m *MessageConstructor) extractPullRequest(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	constr.Write(assets.PullRequest + " ")
	{
		author := payload.GetAuthor()
		constr.WriteWithLink(author.Name, author.Link)
	}
	{
		pr := payload.GetPullRequest()

		switch pr.GetState() {
		case domain.PullRequestStateOpened:

			switch payload.GetAction() {
			case domain.ActionOpened:
				constr.Write(" has opened a pull request ")
				constr.WriteWithLink("\""+pr.Name+"\"", pr.Link)

				constr.Write(" from branch ")
				constr.WriteWithLink(pr.Base.Name, pr.Base.Link)

				constr.Write(" to branch ")
				constr.WriteWithLink(pr.Target.Name, pr.Target.Link)
			case domain.ActionSynchronized:
				constr.Write(" has merged pull request ")
				constr.WriteWithLink("\""+pr.Name+"\"", pr.Link)
			}

		case domain.PullRequestStateClosed:
			constr.Write(" has closed a pull request")
			constr.WriteWithLink("\""+pr.Name+"\"", pr.Link)

		default:
			constr.Write(" has performed something connected to pull request. And we don't know what it is: " + pr.StateStr)
			return constr.String(), constr.format, nil
		}

		commitsAmount := payload.GetCommitsAmount()
		constr.Write(" with " + strconv.Itoa(commitsAmount))
		if commitsAmount%10 == 1 {
			constr.Write(" commit")
		} else {
			constr.Write(" commits")
		}

	}

	return constr.String(), constr.format, nil
}
