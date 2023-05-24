package delete_rule

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
)

type Handler struct {
	rs interfaces.RuleService
}

func (h *Handler) GetCommand() string {
	return commands.DeleteRule
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		rs: srv.RuleService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) == 0 {
		out.SendMessage(constructors.GetEndState("Handler requires 1 argument: id of rule to delete"))
		return
	}

	ruleId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Id has to be positive integer"))
		return
	}

	err = h.rs.DeleteById(context.Background(), ruleId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error obtaining rules: " + err.Error()))
		return
	}

	buttons := &keyboard.InlineKeyboard{}
	buttons.AddButton(assets.Back+"Return to main menu", commands.MainMenu)
	buttons.AddButton(assets.Back+"Return to tickets list", commands.OpenMyTicketsList)

	out.SendMessage(
		constructors.GetEndStateEditMsg(
			"Rule "+strconv.FormatUint(ruleId, 10)+" has been successfully deleted",
			uint64(in.MessageID)))
}

func (h *Handler) GetDescription() string {
	return "Deletes rule with given {{ id }}"
}
