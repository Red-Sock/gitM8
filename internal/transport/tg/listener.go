package tg

import (
	"context"

	"github.com/Red-Sock/go_tg/client"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/main_menu"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/create_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/delete_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/open_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/rename_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/add_rule"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/delete_rule"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/edit_rule"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/open_rule"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/rules/rules_list"
)

type Server struct {
	bot *client.Bot
}

func New(cfg *config.Config, bot *client.Bot, srvs interfaces.Services) (s *Server) {
	s = &Server{}
	s.bot = bot

	{
		s.bot.AddCommandHandler(main_menu.New(srvs))

		s.bot.AddCommandHandler(create_ticket.New(srvs, cfg.GetString(config.WebhookHostURL)))

		s.bot.AddCommandHandler(my_tickets.New(srvs))

		s.bot.AddCommandHandler(open_ticket.New(srvs, cfg.GetString(config.WebhookHostURL)))
		s.bot.AddCommandHandler(rename_ticket.New(srvs))
		s.bot.AddCommandHandler(delete_ticket.New(srvs))

		s.bot.AddCommandHandler(add_rule.New(srvs))
		s.bot.AddCommandHandler(rules_list.New(srvs))
		s.bot.AddCommandHandler(open_rule.New(srvs))
		s.bot.AddCommandHandler(delete_rule.New(srvs))
		s.bot.AddCommandHandler(edit_rule.New(srvs))
	}

	return s
}

func (s *Server) Start(_ context.Context) error {
	return s.bot.Start()
}

func (s *Server) Stop(_ context.Context) error {
	s.bot.Stop()
	return nil
}
