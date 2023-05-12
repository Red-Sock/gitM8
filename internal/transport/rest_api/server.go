package rest_api

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"
)

type Server struct {
	HttpServer *http.Server

	services interfaces.Services

	version string

	crtPath string
	keyPath string
}

func NewServer(cfg *config.Config, services interfaces.Services) *Server {
	r := http.NewServeMux()
	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + cfg.GetString(config.ServerRestAPIPort),
			Handler: r,
		},
		services: services,
		version:  cfg.GetString(config.AppInfoVersion),
		crtPath:  cfg.GetString(config.ServerRestAPICertPath),
		keyPath:  cfg.GetString(config.ServerRestAPIKeyPath),
	}

	r.HandleFunc("/version", s.Version)
	r.HandleFunc("/webhooks/", s.Webhook)
	return s
}

func (s *Server) Start(_ context.Context) error {
	go func() {
		var err error
		if s.crtPath != "" && s.keyPath != "" {
			logrus.Infof("starting webhook https listener on: %s", s.HttpServer.Addr)
			err = s.HttpServer.ListenAndServeTLS(s.crtPath, s.keyPath)
		} else {
			logrus.Infof("starting webhook http listener on: %s", s.HttpServer.Addr)
			err = s.HttpServer.ListenAndServe()
		}
		if err != nil && err != http.ErrServerClosed {
			logrus.Fatal(err)
		}
	}()

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	err := s.HttpServer.Shutdown(ctx)
	if err == http.ErrServerClosed {
		return nil
	}

	return err
}

func (s *Server) formResponse(r interface{}) ([]byte, error) {
	return json.Marshal(r)
}
