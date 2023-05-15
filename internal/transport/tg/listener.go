package tg

import (
	"context"

	"github.com/Red-Sock/go_tg/client"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/create_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/main_menu"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/open_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/open_ticket/delete_ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my_tickets/open_ticket/rename_ticket"
)

type Server struct {
	bot *client.Bot
}

func New(cfg *config.Config, srvs interfaces.Services) (s *Server) {
	s = &Server{}
	s.bot = client.NewBot(cfg.GetString(config.ServerTgAPIKey))

	{
		s.bot.AddCommandHandler(main_menu.New(srvs.TicketsService()), main_menu.Command)

		s.bot.AddCommandHandler(create_ticket.New(srvs.TicketsService()), create_ticket.Command)

		s.bot.AddCommandHandler(my_tickets.New(srvs.TicketsService()), my_tickets.Command)
		s.bot.AddCommandHandler(open_ticket.New(srvs.TicketsService()), open_ticket.Command)
		s.bot.AddCommandHandler(rename_ticket.New(srvs.TicketsService()), rename_ticket.Command)
		s.bot.AddCommandHandler(delete_ticket.New(srvs.TicketsService()), delete_ticket.Command)
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
