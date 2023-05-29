package bootstrap

import (
	"context"

	"github.com/Red-Sock/go_tg/client"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport"
	"github.com/Red-Sock/gitm8/internal/transport/tg"

	"github.com/Red-Sock/gitm8/internal/transport/rest_api"
)

func ApiEntryPoint(
	ctx context.Context,
	cfg *config.Config,
	services interfaces.Services,
	bot *client.Bot,
) (func(context.Context) error, error) {
	mngr := transport.NewManager()

	mngr.AddServer(rest_api.NewServer(cfg, services))
	mngr.AddServer(tg.New(cfg, bot, services))

	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			logrus.Fatalf("error starting server %s", err.Error())
		}
	}()

	return mngr.Stop, nil
}
