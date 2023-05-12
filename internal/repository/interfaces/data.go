package interfaces

import (
	"context"

	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

var (
	// ErrTicketUnavailable - for not existing OR private tickets
	ErrTicketUnavailable = errors.New("ticket does not exists")
)

type Repository interface {
	User() UserRepo
	Ticket() TicketRepo
	Rule() RulesRepo
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
	Get(ctx context.Context, ownerId uint64, url string) (domain.Ticket, error)
	GetById(ctx context.Context, ownerId, id uint64) (domain.Ticket, error)
	GetByUser(ctx context.Context, userID uint64) ([]domain.Ticket, error)
	Rename(ctx context.Context, ownerId, id uint64, newName string) error
	Delete(ctx context.Context, ownerId, id uint64) error
}

type RulesRepo interface {
	Add(ctx context.Context, rule domain.TicketRule) error
	//Get(ctx context.Context) (domain.TicketRule, error)
}
