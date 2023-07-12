package tg_message_constructor

import (
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

func (m *MessageConstructor) extractPushMessage(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	commits := payload.GetCommits()

	author := payload.GetAuthor()
	proj := payload.GetProject()
	srcBranch := payload.GetSrcBranch()

	if len(commits) == 0 {
		constr.Write(assets.Delete)
		constr.WriteWithLink(author.Name, author.Link)

		constr.Write(" has deleted branch ")
		constr.WriteWithLink(srcBranch.Name, srcBranch.Link)

		constr.Write(" at project ")
		constr.WriteWithLink(proj.Name, proj.Link)
		return constr.String(), constr.format, nil
	}

	constr.Write(assets.Push)

	constr.WriteWithLink(author.Name, author.Link)

	constr.Write(" has pushed to project ")
	constr.WriteWithLink(proj.Name, proj.Link)

	constr.Write(" to branch ")
	constr.WriteWithLink(srcBranch.Name, srcBranch.Link)

	constr.Write(" " + strconv.Itoa(len(commits)))
	if len(commits)%10 == 1 {
		constr.Write(" commit")
	} else {
		constr.Write(" commits")
	}

	return constr.String(), constr.format, nil
}
