package bootstrap

import (
	"context"
	"log"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport"
	"github.com/Red-Sock/gitm8/internal/transport/tg"

	"github.com/Red-Sock/gitm8/internal/transport/rest_api"
)

func ApiEntryPoint(ctx context.Context, cfg *config.Config, services interfaces.Services) func(context.Context) error {
	mngr := transport.NewManager()

	mngr.AddServer(rest_api.NewServer(cfg, services))
	mngr.AddServer(tg.New(cfg, services))
	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			log.Fatalf("error starting server %s", err.Error())
		}
	}()
	return mngr.Stop
}
