package tg

import (
	"context"

	"github.com/Red-Sock/go_tg/client"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	create_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/create-ticket"
	my_tickets "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets"
	open_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket"
	delete_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket/delete-ticket"
	rename_ticket "github.com/Red-Sock/gitm8/internal/transport/tg/handlers/my-tickets/open-ticket/rename-ticket"
	"github.com/Red-Sock/gitm8/internal/transport/tg/menus/mainmenu"
)

type Server struct {
	bot *client.Bot
}

func New(cfg *config.Config, srvs interfaces.Services) (s *Server) {
	s = &Server{}
	s.bot = client.NewBot(cfg.GetString(config.ServerTgAPIKey))

	{
		s.bot.AddCommandHandler(create_ticket.New(srvs.TicketsService()), create_ticket.Command)

		s.bot.AddCommandHandler(my_tickets.New(srvs.TicketsService()), my_tickets.Command)
		s.bot.AddCommandHandler(open_ticket.New(srvs.TicketsService()), open_ticket.Command)
		s.bot.AddCommandHandler(rename_ticket.New(srvs.TicketsService()), rename_ticket.Command)
		s.bot.AddCommandHandler(delete_ticket.New(srvs.TicketsService()), delete_ticket.Command)
	}

	{
		s.bot.AddMenu(mainmenu.NewMainMenu())
	}

	return s
}

func (s *Server) Start(_ context.Context) error {
	s.bot.Start()
	return nil
}

func (s *Server) Stop(_ context.Context) error {
	s.bot.Stop()
	return nil
}
