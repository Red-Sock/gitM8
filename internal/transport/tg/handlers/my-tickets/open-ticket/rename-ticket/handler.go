package rename_ticket

import (
	"context"
	"strconv"
	"strings"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/shared_commands"
)

const (
	Command = "/rename-ticket"
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
	if len(in.Args) < 1 {
		out.SendMessage(&response.MessageOut{Text: "Command require 1-2 argument to work properly: id and new name of ticket without spaces or in quotes"})
		return
	}

	ctx := context.Background()
	ticketId, err := strconv.ParseUint(in.Args[0], 10, 10)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "First argument should be positive integer - ticketId"})
		return
	}
	var name string
	if len(in.Args) > 1 {
		name = strings.Join(in.Args[1:], " ")
	} else {
		name, err = h.getNameFromUser(in, out)
		if err != nil {
			out.SendMessage(&response.MessageOut{Text: "Error obtaining new name for ticket from user: " + err.Error()})
			return
		}
	}

	err = h.tickets.Rename(ctx, ticketId, uint64(in.From.ID), name)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error changing name for ticket with id " + in.Args[0] + " to " + in.Args[1]})
		return
	}

	out.SendMessage(response.NewOpenMenu(shared_commands.MainMenu, in))
}

func (h *Handler) getNameFromUser(in *model.MessageIn, out tgapi.Chat) (string, error) {
	out.SendMessage(&response.EditMessage{
		MessageId: int64(in.MessageID),
		Text:      "Entry new name for ticket",
	})

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	rsp, err := out.GetInput(ctx)
	if err != nil {
		return "", err
	}

	out.SendMessage(&response.DeleteMessage{MessageId: int64(rsp.MessageID)})

	return rsp.Text, nil
}
