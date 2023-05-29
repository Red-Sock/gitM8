package rule_constructors

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
)

type Constructor interface {
	Build(out tgapi.Chat) domain.TicketRule

	RecoverFromSource(rule domain.TicketRule) error
}

func SelectConstructor(ruleType domain.RuleType, rs serviceInterfaces.RuleService, messageId, ticketId uint64) (Constructor, error) {
	switch ruleType {
	case domain.RuleTypeWhitelist:
		return &WhitelistConstructor{
			Rs:        rs,
			MessageId: messageId,
			TicketId:  ticketId,
		}, nil
	default:
		return nil, errors.New("Invalid rule type")
	}
}
