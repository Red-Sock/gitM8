package interfaces

import (
	"context"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	"github.com/Red-Sock/gitm8/internal/service/domain/webhook"
)

type Services interface {
	RegistrationService() RegistrationService
	WebhookService() WebhookService
}

type RegistrationService interface {
	CreateBasicTicket(ctx context.Context, request domain.CreateTicketRequest) (domain.Ticket, error)
}

type WebhookService interface {
	HandleWebhook(webhook webhook.Request) error
}
