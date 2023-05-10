package mainmenu

import (
	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/response/menu"

	create_ticket "github.com/Red-Sock/gitM8/internal/transport/tg/handlers/create-ticket"
)

const Command = "/start"

func NewMainMenu() interfaces.Menu {
	m := menu.NewSimple("Main menu", Command)

	m.AddButton("Get url for webhook", create_ticket.Command)

	return m
}
