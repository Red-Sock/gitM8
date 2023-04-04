package pg

import (
	"context"

	"github.com/jackc/pgx/v5"

	"gitM8/internal/service/domain"
)

type TicketRepo struct {
	conn *pgx.Conn
}

func NewTicketRepo(conn *pgx.Conn) *TicketRepo {
	return &TicketRepo{
		conn: conn,
	}
}

func (t *TicketRepo) Add(ctx context.Context, req domain.Ticket) (domain.Ticket, error) {
	err := t.conn.QueryRow(ctx,
		`
INSERT INTO tickets
    (name, owner_id, web_url) VALUES 
	(  $1,       $2,      $3)`,
		req.Name,
		req.WebURL,
		req.OwnerId,
	).Scan(req.Id)
	if err != nil {
		return req, err
	}
	return req, nil
}
