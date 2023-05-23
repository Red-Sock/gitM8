package main_menu

import (
	"context"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
)

type Handler struct {
	tickets serviceInterfaces.TicketsService
}

func (h *Handler) GetCommand() string {
	return commands.MainMenu
}

func New(srv serviceInterfaces.Services) *Handler {
	return &Handler{
		tickets: srv.TicketsService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	menuKeyboard := &keyboard.InlineKeyboard{}
	menuKeyboard.Columns = 2
	menuKeyboard.AddButton("🔗Get url for webhook", commands.CreateTicket)

	tickets, err := h.tickets.GetByUser(context.Background(), uint64(in.From.ID))
	if err != nil {
		out.SendMessage(&response.MessageOut{
			Text: "error obtaining tickets from database: " + err.Error(),
		})
	}

	if len(tickets) > 0 {
		menuKeyboard.AddButton("🎫My tickets", commands.OpenMyTicketsList)
	}

	if in.IsCallback {
		out.SendMessage(&response.EditMessage{
			Text:      "Main menu",
			MessageId: int64(in.MessageID),
			Keys:      menuKeyboard,
		})
	} else {
		out.SendMessage(&response.MessageOut{
			Text: "Main menu",
			Keys: menuKeyboard,
		})
	}

}

func (h *Handler) GetDescription() string {
	return "Opens main menu"
}
