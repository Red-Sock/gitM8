package v1

import (
	"context"

	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type RulesService struct {
	rules dataInterfaces.RulesRepo
}

func NewRuleService(repository dataInterfaces.Repository) *RulesService {
	return &RulesService{
		rules: repository.Rule(),
	}
}

func (w *RulesService) GetRules(ctx context.Context, ticketId uint64) ([]domain.TicketRule, error) {
	tickets, err := w.rules.Get(ctx, ticketId)
	if err != nil {
		return nil, errors.Wrap(err, "error from repository")
	}

	return tickets, nil
}

func (w *RulesService) AddRules(ctx context.Context, rules ...domain.TicketRule) error {
	err := w.rules.Add(ctx, rules...)
	if err != nil {
		return errors.Wrap(err, "error adding rules to database")
	}

	return nil
}
