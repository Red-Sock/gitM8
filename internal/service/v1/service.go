package v1

import (
	"context"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/repository/pg"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/v1/tg-message-constructor"
)

type Service struct {
	regSrv               interfaces.TicketsService
	ruleSrv              interfaces.RuleService
	webhookSrv           interfaces.WebhookService
	messagingConstructor interfaces.MessageConstructor
}

func NewService(ctx context.Context, cfg *config.Config, chat interfaces.Chat) (*Service, error) {
	pgRepo, err := pg.NewRepository(ctx, cfg)
	if err != nil {
		return nil, err
	}

	messagingConstructor := tg_message_constructor.NewMessageConstructor(pgRepo)

	return &Service{
		regSrv:               NewRegistrationService(pgRepo, cfg),
		webhookSrv:           NewWebhookService(pgRepo, messagingConstructor, chat),
		ruleSrv:              NewRuleService(pgRepo),
		messagingConstructor: messagingConstructor,
	}, nil
}

func (s *Service) TicketsService() interfaces.TicketsService {
	return s.regSrv
}

func (s *Service) WebhookService() interfaces.WebhookService {
	return s.webhookSrv
}

func (s *Service) RuleService() interfaces.RuleService {
	return s.ruleSrv
}
