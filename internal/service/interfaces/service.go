package interfaces

import (
	"context"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type Services interface {
	TicketsService() TicketsService
	WebhookService() WebhookService
	RuleService() RuleService
}

type TicketsService interface {
	CreateBasicTicket(ctx context.Context, request domain.CreateTicketRequest) (domain.Ticket, error)
	GetByUser(ctx context.Context, userId uint64) ([]domain.Ticket, error)
	GetById(ctx context.Context, userId, ticketId uint64) (domain.Ticket, error)
	Rename(ctx context.Context, ticketId, userId uint64, newName string) error
	Delete(ctx context.Context, ticketId, userId uint64) error
}

type WebhookService interface {
	HandleWebhook(webhook domain.TicketRequest) error
}

type RuleService interface {
	AddRules(ctx context.Context, rules ...domain.TicketRule) error
	GetRules(ctx context.Context, ticketId uint64) ([]domain.TicketRule, error)
}
