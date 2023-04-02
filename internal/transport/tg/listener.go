package tg

import (
	"github.com/AlexSkilled/go_tg/client"

	"gitM8/internal/config"
	"gitM8/internal/transport/tg/menus/mainmenu"
)

type Server struct {
	bot *client.Bot
}

func New(cfg *config.Config) (s *Server) {
	s = &Server{}
	s.bot = client.NewBot(cfg.GetString(config.Server))

	{
		// handlers
	}

	{
		s.bot.AddMenu(mainmenu.NewMainMenu())
	}

	return s
}

func (s *Server) Start() {
	s.bot.Start()
}

func (s *Server) Stop() {
	s.bot.Stop()
}
