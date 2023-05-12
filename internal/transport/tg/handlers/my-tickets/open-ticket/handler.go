package open_ticket

import (
	"context"
	"fmt"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/Red-Sock/go_tg/model/response/menu"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	delete_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket/delete-ticket"
	rename_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket/rename-ticket"
)

const (
	Command = "/open-ticket"

	ticketInfoPattern = `
Name: %s
Id: %d
URL: %s
GitSystem (auto detected): %s
OwnerId: %d
`
)

type Handler struct {
	tickets interfaces.TicketsService
}

func New(tickets interfaces.TicketsService) *Handler {
	return &Handler{
		tickets: tickets,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) != 1 {
		out.SendMessage(&response.MessageOut{Text: "Command require only 1 argument: id of ticket"})
		return
	}

	ctx := context.Background()

	id, err := strconv.ParseUint(in.Args[0], 10, 10)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Id has to be positive integer"})
		return
	}

	ticket, err := h.tickets.GetById(ctx, uint64(in.From.ID), id)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error obtaining ticket: " + err.Error()})
		return
	}

	url, err := ticket.GetWebUrl()
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error creating web url of ticket: " + err.Error()})
		return
	}

	buttons := &menu.InlineKeyboard{}

	strId := strconv.FormatUint(ticket.Id, 10)
	buttons.AddButton("✍️Rename", rename_ticket.Command+" "+strId)
	buttons.AddButton("🗑️Delete", delete_ticket.Command+" "+strId)

	if ticket.Name == "" {
		ticket.Name = "None"
	}

	out.SendMessage(&response.EditMessage{
		MessageId: int64(in.MessageID),
		Text: fmt.Sprintf(ticketInfoPattern,
			ticket.Name,
			ticket.Id,
			url,
			ticket.GitSystem.String(),
			ticket.OwnerId,
		),
		Keys: buttons,
	})
}
