package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Red-Sock/gitm8/cmd/gitM8/bootstrap"
	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/v1"
"gitM8/cmd/gitM8/bootstrap"
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
stopFunc := bootstrap.ApiEntryPoint(ctx, cfg)

	
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
