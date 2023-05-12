package my_tickets

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/Red-Sock/go_tg/model/response/menu"
	"github.com/sirupsen/logrus"

	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
	open_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket"
)

const Command = "/my-tickets"

type Handler struct {
	tickets serviceInterfaces.TicketsService
}

func New(regService serviceInterfaces.TicketsService) *Handler {
	return &Handler{
		tickets: regService,
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

	buttons := &menu.InlineKeyboard{}

	buttons.Rows = uint8(len(tickets))
	buttons.Columns = 1

	for _, item := range tickets {
		url, err := item.GetWebUrl()
		if err != nil {

			return
		}
		cmd := open_ticket.Command + " " + strconv.FormatUint(item.Id, 10)
		if item.Name != "" {
			buttons.AddButton("🟩"+item.Name, cmd)
		} else {
			buttons.AddButton("🟨"+url, cmd)
		}
	}

	out.SendMessage(&response.EditMessage{
		Text:      "🎫My tickets",
		Keys:      buttons,
		MessageId: int64(in.MessageID),
	})
}
