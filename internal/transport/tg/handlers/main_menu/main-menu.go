package main_menu

import (
	"context"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/Red-Sock/go_tg/model/response/menu"

	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/create_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets"
)

const Command = "/start"

type Handler struct {
	tickets serviceInterfaces.TicketsService
}

func New(regService serviceInterfaces.TicketsService) *Handler {
	return &Handler{
		tickets: regService,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	menuKeyboard := &menu.InlineKeyboard{}
	menuKeyboard.Columns = 2
	menuKeyboard.AddButton("🔗Get url for webhook", create_ticket.Command)

	tickets, err := h.tickets.GetByUser(context.Background(), uint64(in.From.ID))
	if err != nil {
		out.SendMessage(&response.MessageOut{
			Text: "error obtaining tickets from database: " + err.Error(),
		})
	}

	if len(tickets) > 0 {
		menuKeyboard.AddButton("🎫My tickets", my_tickets.Command)
	}

	out.SendMessage(&response.MessageOut{
		Text: "Main menu",
		Keys: menuKeyboard,
	})
}

func (h *Handler) GetDescription() string {
	return "Opens main menu"
}
