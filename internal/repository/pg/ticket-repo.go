package pg

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/service/domain"
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
    (name, owner_id, uri, git_system) VALUES 
	(  $1,       $2,      $3, $4) RETURNING id`,
		req.Name,
		req.OwnerId,
		req.URI,
		req.GitSystem,
	).Scan(&req.Id)
	if err != nil {
		return req, err
	}

	return req, nil
}

func (t *TicketRepo) Get(ctx context.Context, ownerId uint64, uri string) (out domain.Ticket, err error) {
	err = t.conn.QueryRow(ctx,
		`
SELECT 
    id,
	name,
	owner_id,
	uri
FROM tickets
WHERE
	owner_id = $1
AND
    uri = $2
`,
		ownerId,
		uri).
		Scan(
			&out.Id,
			&out.Name,
			&out.OwnerId,
			&out.URI,
		)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return domain.Ticket{}, nil
		}
		return domain.Ticket{}, errors.Wrap(err, "error executing select for ticket via ownerID and uri")
	}

	return out, nil
}

func (t *TicketRepo) GetById(ctx context.Context, ownerId, id uint64) (out domain.Ticket, err error) {
	err = t.conn.QueryRow(ctx,
		`
SELECT 
    id,
	name,
	owner_id,
	uri
FROM tickets
WHERE
	owner_id = $1
AND
    id = $2
`,
		ownerId,
		id).
		Scan(
			&out.Id,
			&out.Name,
			&out.OwnerId,
			&out.URI,
		)

	if err != nil {
		return out, errors.Wrap(err, "error executing select for ticket via ownerID and uri")
	}

	return out, nil
}

func (t *TicketRepo) GetByUser(ctx context.Context, userID uint64) ([]domain.Ticket, error) {
	rows, err := t.conn.Query(ctx,
		`
SELECT 
    id,
	name,
	owner_id,
	uri
FROM tickets
WHERE
	owner_id = $1
`,
		userID)
	defer rows.Close()

	out := make([]domain.Ticket, 0, 1)
	for rows.Next() {
		var tck domain.Ticket
		err = rows.Scan(
			&tck.Id,
			&tck.Name,
			&tck.OwnerId,
			&tck.URI,
		)
		if err != nil {
			return nil, errors.Wrap(err, "error scanning ticket from db")
		}

		out = append(out, tck)
	}

	return out, nil
}

func (t *TicketRepo) Rename(ctx context.Context, userId, ticketId uint64, newName string) error {
	_, err := t.conn.Exec(ctx,
		`
UPDATE tickets 
SET   name     = $1
WHERE id       = $2
AND   owner_id = $3
`,
		newName,
		userId,
		ticketId,
	)
	if err != nil {
		return errors.Wrap(err, "error updating on data layer")
	}
	return nil
}

func (t *TicketRepo) Delete(ctx context.Context, ticketId uint64) error {
	// TODO rethink value of dropping linked tables onto doing so at a service layer
	// TODO in addition to that - this mess down here must exists in order to work properly
	// TODO otherwise. connection block for ever
	batch := &pgx.Batch{}

	batch.Queue(`DELETE FROM ticket_rules  WHERE ticket_id = $1`, ticketId)
	batch.Queue(`DELETE FROM subscriptions WHERE ticket_id = $1`, ticketId)
	batch.Queue(`DELETE FROM tickets       WHERE id        = $1`, ticketId)

	btch := t.conn.SendBatch(ctx, batch)

	defer func() {
		err := btch.Close()
		if err != nil {
			logrus.Errorf("error closing connection after batch while deleting ticket %s", err)
		}
	}()
	_, err := btch.Exec()
	if err != nil {
		return errors.Wrap(err, "error updating on data layer")
	}

	return nil
}

func (t *TicketRepo) HasAccess(ctx context.Context, ticketId, userId uint64) (bool, error) {
	var resp bool

	err := t.conn.QueryRow(ctx, `
SELECT
    EXISTS(
    	SELECT 0 
    	FROM   tickets 
    	WHERE  id = $1 
    	AND    owner_id = $2
    )`,
		ticketId,
		userId,
	).Scan(&resp)
	if err != nil {
		return false, errors.Wrap(err, "error obtaining response from database")
	}

	return resp, nil
}
