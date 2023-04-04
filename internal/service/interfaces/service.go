package interfaces

import (
	"context"

	"gitM8/internal/service/domain"
)

type Services interface {
	RegistrationService() RegistrationService
}

type RegistrationService interface {
	CreateBasicTicket(ctx context.Context, request domain.CreateTicketRequest) (domain.Ticket, error)
}
