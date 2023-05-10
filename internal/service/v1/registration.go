package v1

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"github.com/Red-Sock/gitM8/internal/config"
	dataInterfaces "github.com/Red-Sock/gitM8/internal/repository/interfaces"
	"github.com/Red-Sock/gitM8/internal/service/domain"
)

type RegistrationService struct {
	tickets dataInterfaces.TicketRepo
	user    dataInterfaces.UserRepo
	// getHost - is a function that returns current address where webhook will send info
	getHost func() string
}

func NewRegistrationService(repository dataInterfaces.Repository, cfg *config.Config) *RegistrationService {
	return &RegistrationService{
		tickets: repository.Ticket(),
		user:    repository.User(),
		getHost: func() string {
			return cfg.GetString(config.WebhookHostURL)
		},
	}
}

func (r *RegistrationService) CreateBasicTicket(ctx context.Context, req domain.CreateTicketRequest) (ticket domain.Ticket, err error) {
	user, err := r.user.Upsert(ctx, domain.TgUser{
		TgId: req.OwnerTgId,
	})

	ticket.OwnerId = user.Id
	ticket.WebURL, err = url.JoinPath(r.getHost(), strconv.Itoa(int(req.OwnerTgId)), strconv.FormatInt(time.Now().Unix(), 16))

	ticket, err = r.tickets.Add(ctx, ticket)
	if err != nil {
		return ticket, errors.Wrap(err, "error saving ticket")
	}

	return ticket, nil
}
