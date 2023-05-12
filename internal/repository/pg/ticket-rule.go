package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

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

func (r *TicketRuleRepo) Add(ctx context.Context, req domain.TicketRule) error {
	payload, err := encoder.MarshalToGob(req)
	if err != nil {
		return errors.Wrap(err, "error marshalling ticket rule")
	}

	tags, err := r.conn.Exec(ctx,
		`
			INSERT INTO ticket_rules
			   (ticket_id, payload,     rule_type) VALUES 
			   (       $1,      $2,            $3)`,
		req.GetTicketId(), payload, req.GetType())
	if err != nil {
		return err
	}

	if !tags.Insert() {
		return errors.New("no insert has been performed")
	}

	return nil
}
