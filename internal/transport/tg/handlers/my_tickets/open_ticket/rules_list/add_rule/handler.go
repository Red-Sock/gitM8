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
		out.SendMessage(&response.MessageOut{
			Text: "Adding rule requires 1 argument: id of ticket",
		})
		return
	}

	ticketId, err := strconv.ParseUint(in.Args[0], 10, 64)
	if err != nil {
		out.SendMessage(&response.MessageOut{
			Text: "Ticket id must be positive integer. Got " + in.Args[0],
		})
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
		out.SendMessage(&response.MessageOut{Text: "Error obtaining input from user: " + err.Error()})
		return
	}

	ruleIdx, err := strconv.Atoi(userResponse.Text)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Expected positive integer - index of rule type from given list, got " + userResponse.Text + ". Error parsing: " + err.Error()})
		return
	}
	switch domain.RuleType(ruleIdx) {
	case domain.RuleTypeInvalid:
		out.SendMessage(&response.MessageOut{Text: "Invalid rule type"})
		return
	case domain.RuleTypeWhitelist:
		h.buildWhiteList(uint64(userResponse.MessageID), ticketId, uint64(userResponse.From.ID), out)
		return
	}

}

func (h *Handler) GetDescription() string {
	return "Opens rule creation menu for ticket with {{ id }}"
}

func (h *Handler) buildWhiteList(messageId, ticketId, userId uint64, out tgapi.Chat) {

	msg := &response.EditMessage{
		Text:      "Choose event that will trigger notification",
		MessageId: int64(messageId),
		Keys:      nil,
	}

	out.SendMessage(msg)
}
