package constructors

import (
	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
)

func GetEndState(name string) interfaces.MessageOut {
	buttons := &keyboard.InlineKeyboard{}
	buttons.AddButton(assets.Back+"Return to main menu", commands.MainMenu)
	buttons.AddButton(assets.Back+"Return to tickets list", commands.OpenMyTicketsList)

	return &response.MessageOut{
		Text: name,
		Keys: buttons,
	}
}
