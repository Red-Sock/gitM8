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
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
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
		out.SendMessage(&response.MessageOut{
			Text: "Menu require only 1 argument: id of ticket",
		})
		return
	}
	ticketId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Id has to be positive integer"})
		return
	}
	ctx, cancelF := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancelF()

	rules, err := h.rs.GetRules(ctx, ticketId)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error obtaining rules"})
		return
	}

	var buttons keyboard.InlineKeyboard

	for _, item := range rules {
		buttons.AddButton(item.GetType().String(), item.GetType().String())
	}

	userResponse, err := out.GetInput(ctx)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error obtaining input from user: " + err.Error()})
		return
	}
	_ = userResponse
}

func (h *Handler) GetDescription() string {
	return "Open list of rules for ticket {{ id }}"
}
