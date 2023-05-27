package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type SubscriptionsRepo struct {
	conn *pgx.Conn
}

func NewSubscriptionsRepo(pgConn *pgx.Conn) *SubscriptionsRepo {
	return &SubscriptionsRepo{
		conn: pgConn,
	}
}

func (s *SubscriptionsRepo) AddSubscriber(ctx context.Context, req domain.Subscription) error {
	_, err := s.conn.Exec(ctx, `
INSERT INTO subscriptions 
 	      (chat_id,    user_id,    ticket_id)
VALUES    (     $1,         $2,           $3)`,
		req.ChatId, req.UserId, req.TicketId)
	if err != nil {
		return errors.Wrap(err, "error inserting req into db")
	}

	return nil
}

func (s *SubscriptionsRepo) GetSubscribers(ctx context.Context, ticketId uint64) ([]domain.Subscription, error) {
	rows, err := s.conn.Query(ctx, `
SELECT
	chat_id,  
	user_id,    
	ticket_id
FROM subscriptions
WHERE ticket_id = $1
`, ticketId)
	if err != nil {
		return nil, errors.Wrap(err, "error selecting subsriptions from db")
	}
	defer rows.Close()

	var out []domain.Subscription

	for rows.Next() {
		sub := domain.Subscription{}
		err = rows.Scan(
			&sub.ChatId,
			&sub.UserId,
			&sub.TicketId,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error while scanning subscription from db to service model")
		}
		out = append(out, sub)
	}

	return out, nil
}
