package domain

type RuleType int

const (
	RuleTypeInvalid RuleType = iota
	RuleTypeWhitelist
)

func (rt RuleType) String() string {
	switch rt {
	case RuleTypeWhitelist:
		return "White list"
	default:
		return "Invalid"
	}
}

func GetRulesTypes() []RuleType {
	return []RuleType{
		RuleTypeWhitelist,
	}
}

// TicketRule indicates whether this event should notify user
type TicketRule interface {
	GetId() uint64
	GetTicketId() uint64
	GetType() RuleType
	String() string
	Fire(in TicketRequest) (ok bool)
}

// TicketRuleWhitelist - whitelist for webhook types for ticket
type TicketRuleWhitelist struct {
	Id        uint64      `json:"-"`
	TicketId  uint64      `json:"-"`
	WhiteList []EventType `json:"white_list"`
}

func (rt *TicketRuleWhitelist) Fire(in TicketRequest) bool {
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
func (rt *TicketRuleWhitelist) GetId() uint64 {
	return rt.Id
}

func (rt *TicketRuleWhitelist) GetType() RuleType {
	return RuleTypeWhitelist
}

func (rt *TicketRuleWhitelist) GetTicketId() uint64 {
	return rt.TicketId
}

func (rt *TicketRuleWhitelist) String() string {
	out := "Whitelist on events\n"

	for _, item := range rt.WhiteList {
		out += item.String() + "\n"
	}

	return out
}

// TODO rules for more specific restrictions
// TODO e.g - send message only if or if not specified user is author
// TODO or - send message only if or if not specified branch is target and etc
