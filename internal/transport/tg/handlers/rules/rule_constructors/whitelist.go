package rule_constructors

import (
	"context"
	"fmt"
	"strconv"
	"time"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model/keyboard"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/assets"
)

var ErrCastingTicketRule = errors.New("error casting ticket rule interface to realisation")

const (
	confirmOption     = "confirm"
	allActivityOption = "all activity"
)

type WhitelistConstructor struct {
	Rs        interfaces.RuleService
	MessageId uint64
	TicketId  uint64
	RuleId    uint64
	List      []WhitelistOption
}

type WhitelistOption struct {
	EventType domain.EventType
	Checked   bool
}

func (w *WhitelistConstructor) Build(out tgapi.Chat) domain.TicketRule {
	if len(w.List) == 0 {
		et := domain.GetEventTypes()
		w.List = make([]WhitelistOption, 0, len(et))

		for _, item := range et {
			w.List = append(w.List, WhitelistOption{
				EventType: item,
			})
		}
	}

	w.collectOptions(w.MessageId, out)

	req := &domain.TicketRuleWhitelist{
		TicketId:  w.TicketId,
		WhiteList: make([]domain.EventType, 0, len(w.List)/2),
	}

	for _, item := range w.List {
		if item.Checked {
			req.WhiteList = append(req.WhiteList, item.EventType)
		}
	}

	req.Id = w.RuleId

	return req
}

func (w *WhitelistConstructor) RecoverFromSource(in domain.TicketRule) error {
	whitelist, ok := in.(*domain.TicketRuleWhitelist)
	if !ok {
		return errors.Wrap(ErrCastingTicketRule, fmt.Sprintf("from %T to whitelist", in))
	}

	et := domain.GetEventTypes()
	w.List = make([]WhitelistOption, len(et))

	for idx := range w.List {
		w.List[idx].EventType = et[idx]

		for _, src := range whitelist.WhiteList {
			if src == w.List[idx].EventType {
				w.List[idx].Checked = true
				break
			}
		}
	}
	w.RuleId = whitelist.Id

	return nil
}

func (w *WhitelistConstructor) collectOptions(messageId uint64, out tgapi.Chat) {
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
			// if one item on List is not checked - all should be checked
			// if all items are checked - we should uncheck them
			for _, item := range w.List {
				if !item.Checked {
					newAllState = true
					break
				}
			}

			for idx := range w.List {
				w.List[idx].Checked = newAllState
			}

		default:
			eventIdx, err := strconv.Atoi(resp.Text)
			if err != nil {
				out.SendMessage(&response.MessageOut{Text: "Rule type must be positive integer: " + err.Error()})
				return
			}

			if eventIdx >= len(w.List) || eventIdx < 0 {
				out.SendMessage(&response.MessageOut{Text: "Rule idx from 0 to " + strconv.Itoa(len(w.List))})
				continue
			}

			w.List[eventIdx].Checked = !w.List[eventIdx].Checked
		}

		out.SendMessage(&response.EditMessage{
			Text:      "Choose event that will trigger notification",
			MessageId: int64(messageId),
			Keys:      w.buildKeyboard(),
		})
	}
}

func (w *WhitelistConstructor) buildKeyboard() *keyboard.InlineKeyboard {
	btns := &keyboard.InlineKeyboard{}

	btns.Columns = 2

	for idx, item := range w.List {
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
