package bootstrap

import (
	"context"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/transport"

	"github.com/Red-Sock/gitm8/internal/transport/tg"
	"github.com/Red-Sock/gitm8/internal/transport/rest_api")

func ApiEntryPoint(ctx context.Context, cfg *config.Config) func(context.Context) error {
	mngr := transport.NewManager()

	mngr.AddServer(tg.NewServer(cfg))
	mngr.AddServer(rest_api.NewServer(cfg))
	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			logrus.Fatalf("error starting server %s", err.Error())
		}
	}()
	return mngr.Stop
}
