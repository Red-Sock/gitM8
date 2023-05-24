package edit_rule

import (
	"context"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"

	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/rule_constructors"
)

type Handler struct {
	rs interfaces.RuleService
}

func (h *Handler) GetCommand() string {
	return commands.EditRule
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		rs: srv.RuleService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) == 0 {
		out.SendMessage(constructors.GetEndState("Handler requires 1 argument: id of rule to edit"))
		return
	}

	ruleId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Id has to be positive integer"))
		return
	}

	rule, err := h.rs.GetRuleById(context.Background(), ruleId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error obtaining rules: " + err.Error()))
		return
	}

	if rule == nil {
		if in.IsCallback {
			out.SendMessage(constructors.GetEndStateEditMsg("No rule with id "+in.Args[0], uint64(in.MessageID)))
		} else {
			out.SendMessage(constructors.GetEndState("No rule with id " + in.Args[0]))
		}
		return
	}

	constr, err := rule_constructors.SelectConstructor(rule.GetType(), h.rs, uint64(in.MessageID), rule.GetTicketId())
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error selecting constructor for rule. No such constructor"))
		return
	}

	err = constr.RecoverFromSource(rule)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error recovering from"))
		return
	}

	err = h.rs.UpdateRule(context.Background(), constr.Build(out), uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error updating rule: " + err.Error()))
		return
	}

	out.SendMessage(constructors.GetEndStateEditMsg("Successfully changes rule", uint64(in.MessageID)))

}

func (h *Handler) GetDescription() string {
	return "Edits rule with given {{ id }}"
}
