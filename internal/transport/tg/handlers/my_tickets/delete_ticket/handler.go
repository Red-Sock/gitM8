package delete_ticket

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
)

type Handler struct {
	tickets interfaces.TicketsService
}

func (h *Handler) GetCommand() string {
	return commands.DeleteTicket
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		tickets: srv.TicketsService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) != 1 {
		out.SendMessage(&response.MessageOut{Text: "Command require only 1 argument: id of ticket"})
		return
	}

	ctx := context.Background()
	ticketId, err := strconv.ParseUint(in.Args[0], 10, 10)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Id has to be positive integer"})
		return
	}

	err = h.tickets.Delete(ctx, ticketId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error deleting ticket with id " + in.Args[0] + ": " + err.Error()))
		return
	}

	out.SendMessage(&response.DeleteMessage{
		ChatId:    in.Chat.ID,
		MessageId: int64(in.MessageID),
	})

	out.SendMessage(constructors.GetEndState("Ticket with id: " + strconv.FormatUint(ticketId, 10) + " has been successfully deleted"))
}

func (h *Handler) GetDescription() string {
	return "Deletes ticket with {{ id }}"
}
