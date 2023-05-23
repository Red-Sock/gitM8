package my_tickets

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/sirupsen/logrus"

	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
)

type Handler struct {
	tickets serviceInterfaces.TicketsService
}

func (h *Handler) GetCommand() string {
	return commands.OpenMyTicketsList
}

func New(srv serviceInterfaces.Services) *Handler {
	return &Handler{
		tickets: srv.TicketsService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	ctx := context.Background()

	tickets, err := h.tickets.GetByUser(ctx, uint64(in.From.ID))
	if err != nil {
		logrus.Errorf("error getting tickets: %s", err)
		out.SendMessage(&response.MessageOut{Text: "internal server error: " + err.Error()})
		return
	}

	if len(tickets) == 0 {
		out.SendMessage(&response.MessageOut{
			Text: "No ticket registered",
		})
		return
	}

	buttons := &keyboard.InlineKeyboard{}

	buttons.Rows = uint8(len(tickets))
	buttons.Columns = 1

	for _, item := range tickets {
		url, err := item.GetWebUrl()
		if err != nil {

			return
		}
		cmd := commands.OpenTicketInfo + " " + strconv.FormatUint(item.Id, 10)
		if item.Name != "" {
			buttons.AddButton(assets.GreenSquare+item.Name, cmd)
		} else {
			buttons.AddButton(assets.YellowSquare+url, cmd)
		}
	}

	buttons.AddReturnButton(assets.Back, commands.MainMenu)

	out.SendMessage(&response.EditMessage{
		Text:      "🎫My tickets",
		Keys:      buttons,
		MessageId: int64(in.MessageID),
	})
}

func (h *Handler) GetDescription() string {
	return "Returns all tickets available to you"
}
