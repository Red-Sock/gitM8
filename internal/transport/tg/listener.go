package tg

import (
	"github.com/AlexSkilled/go_tg/client"
	tginterfaces "github.com/AlexSkilled/go_tg/interfaces"
)

type Server struct {
	bot *client.Bot
}

type CreateRequest struct {
	Handlers    map[string]tginterfaces.CommandHandler
	Credentials string
	Menus       []tginterfaces.Menu
}

func New(req CreateRequest) (s *Server) {
	s = &Server{}
	s.bot = client.NewBot(req.Credentials)

	for name, impl := range req.Handlers {
		s.bot.AddCommandHandler(impl, name)
	}

	for _, impl := range req.Menus {
		s.bot.AddMenu(impl)
	}

	return s
}

func (s *Server) Start() {
	s.bot.Start()
}

func (s *Server) Stop() {
	s.bot.Stop()
}
