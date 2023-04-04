package service

import (
	"context"

	"gitM8/internal/config"
	"gitM8/internal/repository/pg"
	"gitM8/internal/service/interfaces"
	v1 "gitM8/internal/service/v1"
)

type Service struct {
	regSrv interfaces.RegistrationService
}

func NewService(ctx context.Context, cfg *config.Config) (*Service, error) {
	pgRepo, err := pg.NewRepository(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &Service{
		regSrv: v1.NewRegistrationService(pgRepo, cfg),
	}, nil
}

func (s *Service) RegistrationService() interfaces.RegistrationService {
	return s.regSrv
}
