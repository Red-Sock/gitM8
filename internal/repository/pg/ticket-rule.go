package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/repository/pg/model"
	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/utils/encoder"
)

type TicketRuleRepo struct {
	conn *pgx.Conn
}

func NewTicketRuleRepo(conn *pgx.Conn) *TicketRuleRepo {
	return &TicketRuleRepo{
		conn: conn,
	}
}

func (r *TicketRuleRepo) Add(ctx context.Context, req ...domain.TicketRule) error {
	dbReq := model.NewTicketRules(req...)
	_, err := r.conn.CopyFrom(ctx,
		dbReq.GetIdentifier(),
		dbReq.GetColumns(),
		dbReq)
	if err != nil {
		return errors.Wrap(err, "error adding ticket rules to database")
	}

	return nil
}

func (r *TicketRuleRepo) Get(ctx context.Context, ticketId uint64) ([]domain.TicketRule, error) {
	rows, err := r.conn.Query(ctx,
		`
SELECT 
	tr.payload,
	tr.rule_type
FROM ticket_rules tr
JOIN public.tickets t on tr.ticket_id = t.id
WHERE
	tr.ticket_id = $1
`,
		ticketId)
	defer rows.Close()

	out := make([]domain.TicketRule, 0, 1)
	for rows.Next() {
		var payload []byte
		var ruleType domain.RuleType
		err = rows.Scan(
			&payload,
			&ruleType,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning ticket from db")
		}

		var tr domain.TicketRule

		switch ruleType {
		case domain.RuleTypeWhitelist:
			tr = &domain.TicketRuleWhitelist{}
		}

		err = encoder.MarshalFromGob(payload, tr)
		if err != nil {
			return nil, errors.Wrap(err, "error marshalling from gob")
		}

		out = append(out, tr)
	}

	return out, nil
}
