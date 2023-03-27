package main

import (
	"context"
	"log"

	"gitM8/internal/config"
	"gitM8/internal/transport"

	"gitM8/internal/transport/rest_api"
)

func apiEntryPoint(ctx context.Context, cfg *config.Config) func(context.Context) error {
	mngr := transport.NewManager()

	mngr.AddServer(rest_api.NewServer(cfg))
	go func() {
		err := mngr.Start(ctx)
		if err != nil {
			log.Fatalf("error starting server %s", err.Error())
		}
	}()
	return mngr.Stop
}
