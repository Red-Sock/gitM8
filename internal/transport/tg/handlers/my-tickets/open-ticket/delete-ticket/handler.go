package delete_ticket

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/shared_commands"
)

const (
	Command = "/delete-ticket"
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
	ticketId, err := strconv.ParseUint(in.Args[0], 10, 10)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Id has to be positive integer"})
		return
	}

	err = h.tickets.Delete(ctx, ticketId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error changing name for ticket with id " + in.Args[0] + " to " + in.Args[1]})
		return
	}

	out.SendMessage(response.NewOpenMenu(shared_commands.MainMenu, in))
}
