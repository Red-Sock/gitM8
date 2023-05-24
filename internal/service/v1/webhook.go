package v1

import (
	"context"

	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
)

type WebhookService struct {
	tickets dataInterfaces.TicketRepo
	rules   dataInterfaces.RulesRepo

	chat serviceInterfaces.Chat
}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (w *WebhookService) HandleWebhook(req domain.TicketRequest) error {
	ctx := context.Background()

	ticket, err := w.tickets.Get(ctx, req.OwnerId, req.Uri)
	if err != nil {
		return errors.Wrap(err, "error from repository")
	}

	if ticket.GitSystem == domain.RepoTypeUnknown {
		ticket.GitSystem = req.Req.Src
	}

	switch ticket.GitSystem {
	case domain.RepoTypeGithub:
		return w.handleGithub(ctx, req, ticket)
	}

	return nil
}

func (w *WebhookService) handleGithub(ctx context.Context, req domain.TicketRequest, ticket domain.Ticket) error {
	rules, err := w.rules.GetByTicketId(ctx, ticket.Id, req.OwnerId)
	if err != nil {
		return errors.Wrap(err, "error obtaining rules")
	}

	_ = rules

	//w.chat.Send()
	return nil
}
