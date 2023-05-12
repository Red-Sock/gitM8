package mainmenu

import (
	"github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/response/menu"

	create_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/create-ticket"
	my_tickets "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/shared_commands"
)

func NewMainMenu() interfaces.Menu {
	m := menu.NewSimple("Main menu", shared_commands.MainMenu)
	m.SetColumnsAmount(2)
	m.AddButton("🔗Get url for webhook", create_ticket.Command)
	m.AddButton("🎫My tickets", my_tickets.Command)
	return m
}
