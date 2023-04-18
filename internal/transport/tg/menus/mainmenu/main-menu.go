package mainmenu

import (
	"github.com/AlexSkilled/go_tg/interfaces"
	"github.com/AlexSkilled/go_tg/model/response/menu"

	create_ticket "gitM8/internal/transport/tg/handlers/create-ticket"
)

const Command = "/start"

func NewMainMenu() interfaces.Menu {
	m := menu.NewSimple("Main menu", Command)

	m.AddButton("Get url for webhook", create_ticket.Command)

	return m
}
