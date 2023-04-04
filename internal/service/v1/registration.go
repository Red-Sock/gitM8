package v1

import (
	"context"
	"net/url"
	"strconv"
	"time"

	"github.com/pkg/errors"

	"gitM8/internal/config"
	dataInterfaces "gitM8/internal/repository/interfaces"
	"gitM8/internal/service/domain"
)

type RegistrationService struct {
	tickets dataInterfaces.TicketRepo
	// getHost - is a function that returns current address where webhook will send info
	getHost func() string
}

func NewRegistrationService(repository dataInterfaces.Repository, cfg *config.Config) *RegistrationService {
	return &RegistrationService{
		tickets: repository.Ticket(),
		//getHost: func() string {
		//	return cfg.GetString(config.)
		//},
	}
}

func (r *RegistrationService) CreateBasicTicket(ctx context.Context, req domain.CreateTicketRequest) (ticket domain.Ticket, err error) {
	ticket.OwnerId = req.OwnerId
	ticket.WebURL, err = url.JoinPath(r.getHost(), strconv.Itoa(int(req.OwnerId)), strconv.FormatInt(time.Now().Unix(), 16))

	ticket, err = r.tickets.Add(ctx, ticket)
	if err != nil {
		return ticket, errors.Wrap(err, "error saving ticket")
	}
	return ticket, nil
}
