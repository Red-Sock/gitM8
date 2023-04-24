package v1

import (
	"context"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/repository/pg"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
)

type Service struct {
	regSrv     interfaces.RegistrationService
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
	}, nil
}

func (s *Service) RegistrationService() interfaces.RegistrationService {
	return s.regSrv
}

func (s *Service) WebhookService() interfaces.WebhookService {
	return s.webhookSrv
}
