package interfaces

import (
	"context"

	"gitM8/internal/service/domain"
	"gitM8/internal/service/domain/webhook"
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
