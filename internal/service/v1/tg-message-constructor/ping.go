package tg_message_constructor

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

func (m *MessageConstructor) extractPingMessage(payload domain.Payload) (string, []tgbotapi.MessageEntity, error) {
	constr := constructor{}

	constr.Write(assets.Ping)

	constr.Write("Repository ")
	{
		proj := payload.GetProject()
		constr.WriteWithLink(proj.Name, proj.Link)
	}
	constr.Write(" has pinged this webhook!")
	constr.Writeln("Sending a pong right away")

	return constr.String(), constr.format, nil
}
