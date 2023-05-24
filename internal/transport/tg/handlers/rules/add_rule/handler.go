package add_rule

import (
	"context"
	"strconv"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/rule_constructors"
)

type Handler struct {
	rs interfaces.RuleService
}

func (h *Handler) GetCommand() string {
	return commands.AddRule
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		rs: srv.RuleService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) == 0 {
		out.SendMessage(constructors.GetEndState("Adding rule requires 1 argument: id of ticket"))
		return
	}

	ticketId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Ticket id must be positive integer. Got " + in.Args[0]))
		return
	}

	ctx, cancelF := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancelF()

	rules := domain.GetRulesTypes()
	var buttons keyboard.InlineKeyboard

	for _, item := range rules {
		buttons.AddButton(item.String(), strconv.Itoa(int(item)))
	}

	out.SendMessage(&response.EditMessage{
		Text:      "Select rule type",
		MessageId: int64(in.MessageID),
		Keys:      &buttons,
	})

	userResponse, err := out.GetInput(ctx)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error obtaining input from user: " + err.Error()))
		return
	}

	ruleIdx, err := strconv.Atoi(userResponse.Text)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Expected positive integer - index of rule type from given list, got " + userResponse.Text + ". Error parsing: " + err.Error()))
		return
	}

	constr, err := rule_constructors.SelectConstructor(domain.RuleType(ruleIdx), h.rs, uint64(in.MessageID), ticketId)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error selecting constructor for rule. No such constructor"))
		return
	}

	err = h.rs.AddRules(context.Background(), constr.Build(out))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error saving rule for ticket: " + err.Error()))
		return
	}

	out.SendMessage(constructors.GetEndState("Successfully created a whitelist rule"))
}

func (h *Handler) GetDescription() string {
	return "Opens rule creation menu for ticket with {{ id }}"
}
