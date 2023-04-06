package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"gitM8/cmd/gitM8/bootstrap"
	"gitM8/internal/config"
	"gitM8/internal/service/v1"
)

func main() {
	log.Println("starting app")

	ctx := context.Background()

	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatalf("error reading config %s", err.Error())
	}

	startupDuration, err := cfg.GetDuration(config.AppInfoStartupDuration)
	if err != nil {
		log.Fatalf("error extracting startup duration %s", err)
	}
	context.WithTimeout(ctx, startupDuration)

	srv, err := v1.NewService(ctx, cfg)
	if err != nil {
		log.Fatalf("error assembling service layer %s", err)
	}

	stopFunc := bootstrap.ApiEntryPoint(ctx, cfg, srv)

	waitingForTheEnd()

	err = stopFunc(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	log.Println("shutting down the app")
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is a LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	<-done
}
