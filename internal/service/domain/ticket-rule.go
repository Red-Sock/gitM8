package domain

type ruleType int

const (
	RuleTypeWhitelist ruleType = iota
)

// TicketRule indicates whether this event should notify user
type TicketRule interface {
	GetId() uint64
	GetTicketId() uint64
	GetType() ruleType
	Fire(in TicketRequest) (ok bool)
}

// RestrictingTicket - whitelist for webhook types for ticket
type RestrictingTicket struct {
	Id        uint64
	TicketId  uint64
	WhiteList []Type
}

func (rt *RestrictingTicket) Fire(in TicketRequest) bool {
	if len(rt.WhiteList) == 0 {
		return true
	}

	for _, wt := range rt.WhiteList {
		if wt == in.Req.Type {
			return true
		}
	}

	return false
}
func (rt *RestrictingTicket) GetId() uint64 {
	return rt.Id
}

func (rt *RestrictingTicket) GetType() ruleType {
	return RuleTypeWhitelist
}

func (rt *RestrictingTicket) GetTicketId() uint64 {
	return rt.TicketId
}

// TODO rules for more specific restrictions
// TODO e.g - send message only if or if not specified user is author
// TODO or - send message only if or if not specified branch is target and etc
