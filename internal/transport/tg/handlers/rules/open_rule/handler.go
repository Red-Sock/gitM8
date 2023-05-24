package open_rule

import (
	"context"
	"fmt"
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

const fmtRuleInfo = `
Rule info for: %d
Type: %s
Description: 
%s
`

type Handler struct {
	rs interfaces.RuleService
}

func (h *Handler) GetCommand() string {
	return commands.OpenRule
}

func New(srv interfaces.Services) *Handler {
	return &Handler{
		rs: srv.RuleService(),
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	if len(in.Args) == 0 {
		out.SendMessage(constructors.GetEndState("Handler requires 1 argument: id of rule"))
		return
	}
	ruleId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(constructors.GetEndState("Id has to be positive integer"))
		return
	}
	ctx, cancelF := context.WithTimeout(context.Background(), time.Minute*1)
	defer cancelF()

	rule, err := h.rs.GetRuleById(ctx, ruleId, uint64(in.From.ID))
	if err != nil {
		out.SendMessage(constructors.GetEndState("Error obtaining rules: " + err.Error()))
		return
	}

	buttons := &keyboard.InlineKeyboard{}

	buttons.AddButton(
		assets.Edit+"Edit",
		commands.EditRule+" "+in.Args[0])
	buttons.AddButton(
		assets.Trash+"Delete",
		commands.DeleteRule+" "+in.Args[0])
	buttons.AddButton(
		assets.Back+"Back",
		commands.OpenRulesList+" "+strconv.FormatUint(rule.GetTicketId(), 10))

	out.SendMessage(
		&response.EditMessage{
			Text:      fmt.Sprintf(fmtRuleInfo, ruleId, rule.GetType(), rule.String()),
			MessageId: int64(in.MessageID),
			Keys:      buttons,
		})
}

func (h *Handler) GetDescription() string {
	return "Open information on rule with given {{ id }}"
}
