package open_ticket

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
)

const (
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
	rules   interfaces.RuleService

	host string
}

func (h *Handler) GetCommand() string {
	return commands.OpenTicketInfo
}

func New(servs interfaces.Services, host string) *Handler {
	return &Handler{
		tickets: servs.TicketsService(),
		rules:   servs.RuleService(),
		host:    host,
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

	webUrl, err := ticket.GetWebUrl()
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error creating web url of ticket: " + err.Error()})
		return
	}
	webUrl, err = url.JoinPath(h.host, webUrl)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error concatenating web url for ticket: " + err.Error()})
		return
	}

	buttons := &keyboard.InlineKeyboard{}

	strId := strconv.FormatUint(ticket.Id, 10)
	buttons.AddButton("✍️Rename", commands.RenameTicket+" "+strId)

	rules, err := h.rules.GetRules(ctx, ticket.Id)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error obtaining rules from db: " + err.Error()})
		return
	}
	if len(rules) != 0 {
		buttons.AddButton("📔Rules", commands.OpenRulesList+" "+strId)
	} else {
		buttons.AddButton("📝Add rule", commands.AddRule+" "+strId)
	}

	buttons.AddButton("🗑️Delete", commands.DeleteTicket+" "+strId)

	buttons.AddReturnButton(assets.Back, commands.OpenMyTicketsList)
	if ticket.Name == "" {
		ticket.Name = "None"
	}

	out.SendMessage(&response.EditMessage{
		MessageId: int64(in.MessageID),
		Text: fmt.Sprintf(ticketInfoPattern,
			ticket.Name,
			ticket.Id,
			webUrl,
			ticket.GitSystem.String(),
			ticket.OwnerId,
		),
		Keys: buttons,
	})
}

func (h *Handler) GetDescription() string {
	return "Returns information about ticket with {{ id }}"
}
