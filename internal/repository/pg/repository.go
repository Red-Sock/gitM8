package pg

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Red-Sock/gitM8/internal/clients/postgres"
	"github.com/Red-Sock/gitM8/internal/config"
	"github.com/Red-Sock/gitM8/internal/repository/interfaces"
)

type Repository struct {
	user   interfaces.UserRepo
	ticket interfaces.TicketRepo
}

func NewRepository(ctx context.Context, cfg *config.Config) (*Repository, error) {
	pgConn, err := postgres.New(ctx, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "couldn't establish connection with postgres")
	}

	return &Repository{
		user:   NewTgUserRepo(pgConn),
		ticket: NewTicketRepo(pgConn),
	}, nil
}

func (r *Repository) User() interfaces.UserRepo {
	return r.user
}

func (r *Repository) Ticket() interfaces.TicketRepo {
	return r.ticket
}
