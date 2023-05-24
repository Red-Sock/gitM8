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

func (t *TicketRuleRepo) AddRules(ctx context.Context, req ...domain.TicketRule) error {
	dbReq := model.NewTicketRules(req...)
	_, err := t.conn.CopyFrom(ctx,
		dbReq.GetIdentifier(),
		dbReq.GetColumns(),
		dbReq)
	if err != nil {
		return errors.Wrap(err, "error adding ticket rules to database")
	}

	return nil
}

func (t *TicketRuleRepo) GetByTicketId(ctx context.Context, ticketId, userId uint64) ([]domain.TicketRule, error) {
	rows, err := t.conn.Query(ctx,
		`
SELECT 
    tr.id,
	tr.payload,
	tr.rule_type
FROM ticket_rules tr
JOIN public.tickets t on tr.ticket_id = t.id AND t.owner_id = $1
WHERE
	tr.ticket_id = $2
`,
		userId, ticketId)
	defer rows.Close()

	out := make([]domain.TicketRule, 0, 1)
	for rows.Next() {
		var payload []byte
		var ruleType domain.RuleType
		var id uint64
		err = rows.Scan(
			&id,
			&payload,
			&ruleType,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning ticket from db")
		}

		var tr domain.TicketRule

		switch ruleType {
		case domain.RuleTypeWhitelist:
			tr = &domain.TicketRuleWhitelist{
				Id: id,
			}
		}

		err = encoder.MarshalFrom(payload, tr)
		if err != nil {
			return nil, errors.Wrap(err, "error marshalling from gob")
		}

		out = append(out, tr)
	}

	return out, nil
}

func (t *TicketRuleRepo) GetById(ctx context.Context, ruleId, userId uint64) (domain.TicketRule, error) {
	var payload []byte
	var ruleType domain.RuleType
	var ticketId uint64

	err := t.conn.QueryRow(ctx,
		`
SELECT 
    tr.ticket_id,
	tr.payload,
	tr.rule_type
FROM ticket_rules tr
JOIN public.tickets t ON tr.ticket_id = t.id 
AND t.owner_id = $1
WHERE
	tr.id = $2
`,
		userId, ruleId).Scan(
		&ticketId,
		&payload,
		&ruleType,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, errors.Wrap(err, "error obtaining from db")
	}
	var rule domain.TicketRule

	switch ruleType {
	case domain.RuleTypeWhitelist:
		rule = &domain.TicketRuleWhitelist{
			Id:       ruleId,
			TicketId: ticketId,
		}
	}

	err = encoder.MarshalFrom(payload, rule)
	if err != nil {
		return nil, errors.Wrap(err, "error marshalling from gob")
	}

	return rule, nil
}

func (t *TicketRuleRepo) DeleteById(ctx context.Context, ruleId, userId uint64) error {
	_, err := t.conn.Exec(ctx, `
DELETE FROM ticket_rules 
WHERE id = $1 
 	AND 
EXISTS( 
	SELECT t.id FROM tickets as t 
	JOIN ticket_rules AS tr ON  tr.ticket_id = t.id
	WHERE tr.id      = $1
	AND   t.owner_id = $2
	          )`,
		ruleId, userId)

	if err != nil {
		return errors.Wrap(err, "error deleting from database")
	}

	return nil
}

func (t *TicketRuleRepo) UpdateRule(ctx context.Context, req domain.TicketRule) error {
	payload, err := encoder.MarshalTo(req)
	if err != nil {
		return errors.Wrap(err, "error marshalling rule to binary")
	}

	_, err = t.conn.Exec(ctx,
		`
	UPDATE ticket_rules 
	SET    payload = $2
	WHERE  id = $1
`,
		req.GetId(),
		payload,
	)
	if err != nil {
		return errors.Wrap(err, "error adding ticket rules to database")
	}

	return nil
}
