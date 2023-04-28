package interfaces

import (
	"context"

	"github.com/Red-Sock/gitM8/internal/service/domain"
)

type Repository interface {
	User() UserRepo
	Ticket() TicketRepo
}

type UserRepo interface {
	// Upsert creates new user and returns his id in database
	// in case if user exists - simply returns his id in database
	Upsert(ctx context.Context, user domain.TgUser) (domain.TgUser, error)

	// Get returns information about user with given telegram_id
	Get(ctx context.Context, tgId int64) (domain.TgUser, error)
}

type TicketRepo interface {
	Add(ctx context.Context, req domain.Ticket) (domain.Ticket, error)
}
