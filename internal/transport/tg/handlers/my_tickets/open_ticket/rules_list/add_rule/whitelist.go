package add_rule

import (
	"context"
	"strconv"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

const (
	confirmOption     = "confirm"
	allActivityOption = "all activity"
)

type whitelistOption struct {
	EventType domain.EventType
	Checked   bool
}

type whitelistConstructor struct {
	list []whitelistOption
}

func (h *Handler) buildWhiteList(messageId, ticketId uint64, out tgapi.Chat) {
	et := domain.GetEventTypes()

	constructor := whitelistConstructor{
		list: make([]whitelistOption, 0, len(et)),
	}

	for _, item := range et {
		constructor.list = append(constructor.list, whitelistOption{
			EventType: item,
		})
	}

	constructor.collectOptions(messageId, out)

	req := &domain.TicketRuleWhitelist{
		TicketId:  ticketId,
		WhiteList: make([]domain.EventType, 0, len(constructor.list)/2),
	}

	for _, item := range constructor.list {
		if item.Checked {
			req.WhiteList = append(req.WhiteList, item.EventType)
		}
	}

	err := h.rs.AddRules(context.Background(), req)
	if err != nil {
		out.SendMessage(&response.MessageOut{Text: "Error saving rule for ticket: " + err.Error()})
		return
	}
}

func (w *whitelistConstructor) collectOptions(messageId uint64, out tgapi.Chat) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	defer out.SendMessage(&response.DeleteMessage{
		MessageId: int64(messageId),
	})

	out.SendMessage(&response.EditMessage{
		Text:      "Choose event that will trigger notification",
		MessageId: int64(messageId),
		Keys:      w.buildKeyboard(),
	})

	for {
		resp, err := out.GetInput(ctx)
		if err != nil {
			out.SendMessage(&response.MessageOut{
				Text: "Error configuring whitelist rule: " + err.Error(),
			})
			return
		}

		switch resp.Text {
		case confirmOption:
			return
		case allActivityOption:
			newAllState := false
			// if one item on list is not checked - all should be checked
			// if all items are checked - we should uncheck them
			for _, item := range w.list {
				if !item.Checked {
					newAllState = true
					break
				}
			}

			for idx := range w.list {
				w.list[idx].Checked = newAllState
			}

		default:
			eventIdx, err := strconv.Atoi(resp.Text)
			if err != nil {
				out.SendMessage(&response.MessageOut{Text: "Rule type must be positive integer: " + err.Error()})
				return
			}

			if eventIdx >= len(w.list) || eventIdx < 0 {
				out.SendMessage(&response.MessageOut{Text: "Rule idx from 0 to " + strconv.Itoa(len(w.list))})
				continue
			}

			w.list[eventIdx].Checked = !w.list[eventIdx].Checked
		}

		out.SendMessage(&response.EditMessage{
			Text:      "Choose event that will trigger notification",
			MessageId: int64(messageId),
			Keys:      w.buildKeyboard(),
		})
	}
}

func (w *whitelistConstructor) buildKeyboard() *keyboard.InlineKeyboard {
	btns := &keyboard.InlineKeyboard{}

	btns.Columns = 2

	for idx, item := range w.list {
		command := strconv.Itoa(idx)

		if item.Checked {
			btns.AddButton(item.EventType.String()+assets.Checked, command)
		} else {
			btns.AddButton(item.EventType.String(), command)
		}
	}

	btns.AddButton(confirmOption+assets.Confirm, confirmOption)
	btns.AddButton(allActivityOption+assets.Confirm, allActivityOption)
	return btns
}
