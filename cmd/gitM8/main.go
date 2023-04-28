package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitM8/cmd/gitM8/bootstrap"
	"github.com/Red-Sock/gitM8/internal/config"
	"github.com/Red-Sock/gitM8/internal/service/v1"
)

func main() {

	ctx := context.Background()

	logrus.Info("reading config")
	cfg, err := config.ReadConfig()
	if err != nil {
		logrus.Fatalf("error reading config %s", err.Error())
	}

	startupDuration, err := cfg.GetDuration(config.AppInfoStartupDuration)
	if err != nil {
		logrus.Fatalf("error extracting startup duration %s", err)
	}

	logrus.Infof("time on startup: %v m %v s", startupDuration.Minutes(), startupDuration.Minutes()*60-startupDuration.Seconds())

	ctx, _ = context.WithTimeout(ctx, startupDuration)

	logrus.Info("initializing service layer")
	srv, err := v1.NewService(ctx, cfg)
	if err != nil {
		logrus.Fatalf("error assembling service layer %s", err)
	}

	logrus.Info("bootstrapping api")
	stopFunc, err := bootstrap.ApiEntryPoint(ctx, cfg, srv)
	if err != nil {
		logrus.Fatalf("error starting api %s", err)
	}

	waitingForTheEnd()

	err = stopFunc(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	err = stopFunc(context.Background())
	if err != nil {
		logrus.Fatal(err)
	}
	logrus.Println("app is shut down")
}

// rscli comment: an obligatory function for tool to work properly.
// must be called in the main function above
// also this is a LP song name reference, so no rules can be applied to the function name
func waitingForTheEnd() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-done

	logrus.Infof("%s signal received, gracefully shutting down", sig)
}
