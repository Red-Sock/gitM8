package bootstrap

import (
	"context"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitM8/internal/config"
	"github.com/Red-Sock/gitM8/internal/service/interfaces"
	"github.com/Red-Sock/gitM8/internal/transport"
	"github.com/Red-Sock/gitM8/internal/transport/rest_api"
	"github.com/Red-Sock/gitM8/internal/transport/tg"
)

func ApiEntryPoint(ctx context.Context, cfg *config.Config, services interfaces.Services) (func(context.Context) error, error) {
	mngr := transport.NewManager()

	srv, err := rest_api.NewServer(cfg, services)
	if err != nil {
		return nil, errors.Wrap(err, "error starting up a server")
	}
	mngr.AddServer(srv)

	mngr.AddServer(tg.New(cfg, services))

	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			logrus.Fatalf("error starting server %s", err.Error())
		}
	}()

	return mngr.Stop, nil
}
