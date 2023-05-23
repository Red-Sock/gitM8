package v1

import (
	"context"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/repository/pg"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
)

type Service struct {
	regSrv     interfaces.TicketsService
	ruleSrv    interfaces.RuleService
	webhookSrv interfaces.WebhookService
}

func NewService(ctx context.Context, cfg *config.Config) (*Service, error) {
	pgRepo, err := pg.NewRepository(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		regSrv:     NewRegistrationService(pgRepo, cfg),
		webhookSrv: NewWebhookService(),
		ruleSrv:    NewRuleService(pgRepo),
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
