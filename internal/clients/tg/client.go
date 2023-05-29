package tg

import (
	"github.com/Red-Sock/go_tg/client"

	"github.com/Red-Sock/gitm8/internal/config"
)

func New(cfg *config.Config) *client.Bot {
	return client.NewBot(cfg.GetString(config.ServerTgAPIKey))
}
