package v1

import (
	"context"

	"github.com/pkg/errors"

	dataInterfaces "github.com/Red-Sock/gitm8/internal/repository/interfaces"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

type RulesService struct {
	rules   dataInterfaces.RulesRepo
	tickets dataInterfaces.TicketRepo
}

func NewRuleService(repository dataInterfaces.Repository) *RulesService {
	return &RulesService{
		rules:   repository.Rule(),
		tickets: repository.Ticket(),
	}
}

func (w *RulesService) GetRulesByTicketId(ctx context.Context, ticketId, userId uint64) ([]domain.TicketRule, error) {
	rules, err := w.rules.GetByTicketId(ctx, ticketId, userId)
	if err != nil {
		return nil, errors.Wrap(err, "error from repository")
	}

	return rules, nil
}

func (w *RulesService) GetRuleById(ctx context.Context, ruleId, userId uint64) (domain.TicketRule, error) {
	rule, err := w.rules.GetById(ctx, ruleId, userId)
	if err != nil {
		return nil, errors.Wrap(err, "error from repository")
	}

	return rule, nil
}

func (w *RulesService) AddRules(ctx context.Context, rules ...domain.TicketRule) error {
	err := w.rules.AddRules(ctx, rules...)
	if err != nil {
		return errors.Wrap(err, "error adding rules to repository")
	}

	return nil
}

func (w *RulesService) DeleteById(ctx context.Context, ruleId, userId uint64) error {
	err := w.rules.DeleteById(ctx, ruleId, userId)
	if err != nil {
		return errors.Wrap(err, "error deleting from repository")
	}

	return nil

}

func (w *RulesService) UpdateRule(ctx context.Context, rule domain.TicketRule, userId uint64) error {
	hasAccess, err := w.tickets.HasAccess(ctx, rule.GetTicketId(), userId)
	if err != nil {
		return errors.Wrap(err, "error obtaining access value")
	}
	if !hasAccess {
		return dataInterfaces.ErrTicketUnavailable
	}

	err = w.rules.UpdateRule(ctx, rule)
	if err != nil {
		return errors.Wrap(err, "error updating rule in repository")
	}
	return nil
}
