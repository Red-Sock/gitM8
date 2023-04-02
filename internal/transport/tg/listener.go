package tg

import (
	"context"

	"github.com/AlexSkilled/go_tg/client"

	"gitM8/internal/config"
	"gitM8/internal/transport/tg/handlers/register"
	"gitM8/internal/transport/tg/menus/mainmenu"
)

type Server struct {
	bot *client.Bot
}

func New(cfg *config.Config) (s *Server) {
	s = &Server{}
	s.bot = client.NewBot(cfg.GetString(config.ServerTgApiKey))

	{
		s.bot.AddCommandHandler(register.New(), register.Command)
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
