package rules_list

import (
	"context"
	"strconv"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
)

type Handler struct {
	rs interfaces.RuleService
}

func (h *Handler) GetCommand() string {
	return commands.OpenRulesList
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		rs: srv.RuleService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) == 0 {
		out.SendMessage(constructors.GetEndState("Menu require only 1 argument: id of ticket"))
		return
	}
	ticketId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Id has to be positive integer"))
		return
	}
	ctx, cancelF := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancelF()

	rules, err := h.rs.GetRulesByTicketId(ctx, ticketId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error obtaining rules " + err.Error()))
		return
	}

	buttons := &keyboard.InlineKeyboard{}

	for _, item := range rules {
		buttons.AddButton(
			item.GetType().String(),
			commands.OpenRule+" "+strconv.FormatUint(item.GetId(), 10))
	}

	buttons.AddButton(assets.Back, commands.OpenTicketInfo+" "+strconv.FormatUint(rules[0].GetTicketId(), 10))

	out.SendMessage(&response.EditMessage{
		Text:      "Rules for ticket " + strconv.FormatUint(ticketId, 10),
		MessageId: int64(in.MessageID),
		Keys:      buttons,
	})
}

func (h *Handler) GetDescription() string {
	return "Open list of rules for ticket {{ id }}"
}
