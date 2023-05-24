package model

import (
	stderr "errors"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/utils/encoder"
)

type TicketRules struct {
	idx   int
	rules []domain.TicketRule
	err   error
}

func NewTicketRules(rules ...domain.TicketRule) *TicketRules {
	return &TicketRules{
		idx:   -1,
		rules: rules,
	}
}

func (tr *TicketRules) Next() bool {
	tr.idx++
	return tr.idx < len(tr.rules)
}

func (tr *TicketRules) Values() ([]any, error) {
	payload, err := encoder.MarshalTo(tr.rules[tr.idx])
	if err != nil {
		tr.err = stderr.Join(tr.err, err)
		return nil, errors.Wrap(err, "error marshalling ticket rule")
	}

	var values = []any{
		tr.rules[tr.idx].GetTicketId(),
		payload,
		tr.rules[tr.idx].GetType(),
	}

	return values, nil
}

func (tr *TicketRules) Err() error {
	return tr.err
}

func (tr *TicketRules) GetIdentifier() pgx.Identifier {
	return []string{"ticket_rules"}
}

func (tr *TicketRules) GetColumns() []string {
	return []string{"ticket_id", "payload", "rule_type"}
}
