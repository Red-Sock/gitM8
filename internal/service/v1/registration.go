package v1

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/Red-Sock/gitm8/internal/config"
	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type TicketService struct {
	tickets dataInterfaces.TicketRepo
	user    dataInterfaces.UserRepo
	// getHost - is a function that returns current address where webhook will send info
	getHost func() string
}

func NewRegistrationService(repository dataInterfaces.Repository, cfg *config.Config) *TicketService {
	return &TicketService{
		tickets: repository.Ticket(),
		user:    repository.User(),
		getHost: func() string {
			return cfg.GetString(config.WebhookHostURL)
		},
	}
}

func (r *TicketService) CreateBasicTicket(ctx context.Context, req domain.CreateTicketRequest) (ticket domain.Ticket, err error) {
	user, err := r.user.Upsert(ctx, domain.TgUser{
		TgId: req.OwnerTgId,
	})

	ticket.OwnerId = user.TgId
	ticket.URI = strconv.FormatInt(time.Now().Unix(), 16)

	ticket, err = r.tickets.Add(ctx, ticket)
	if err != nil {
		return ticket, errors.Wrap(err, "error saving ticket")
	}

	return ticket, nil
}

func (r *TicketService) GetByUser(ctx context.Context, userId uint64) ([]domain.Ticket, error) {
	tickets, err := r.tickets.GetByUser(ctx, userId)
	if err != nil {
		return nil, errors.Wrap(err, "error obtaining tickets for user from storage")
	}

	return tickets, nil
}

func (r *TicketService) GetById(ctx context.Context, userId, ticketId uint64) (domain.Ticket, error) {
	ticket, err := r.tickets.GetById(ctx, userId, ticketId)
	if err != nil {
		return domain.Ticket{}, errors.Wrap(err, "error obtaining ticket from storage")
	}

	return ticket, nil
}

func (r *TicketService) Rename(ctx context.Context, ticketId, userId uint64, newName string) error {
	return r.tickets.Rename(ctx, ticketId, userId, newName)
}

func (r *TicketService) Delete(ctx context.Context, ticketId, userId uint64) error {
	return r.tickets.Delete(ctx, ticketId, userId)
}
