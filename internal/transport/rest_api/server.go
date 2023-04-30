package rest_api

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitM8/internal/config"
	"github.com/Red-Sock/gitM8/internal/service/interfaces"

	"github.com/gorilla/mux"
)

type Server struct {
	HttpServer *http.Server

	services interfaces.Services

	version string

	crtPath string
	keyPath string
}

func NewServer(cfg *config.Config, services interfaces.Services) (*Server, error) {
	port, err := cfg.TryGetInt(config.ServerRestAPIPort)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting "+config.ServerRestAPIPort+" from config")
	}

	r := mux.NewRouter()
	s := &Server{
		HttpServer: &http.Server{
			Addr:    ":" + strconv.Itoa(port),
			Handler: r,
		},
		services: services,
		version:  cfg.GetString(config.AppInfoVersion),
		crtPath:  cfg.GetString(config.ServerRestAPICertPath),
		keyPath:  cfg.GetString(config.ServerRestAPIKeyPath),
	}

	r.HandleFunc("/version", s.Version)
	r.HandleFunc("/webhooks", s.Webhook)
	return s, nil
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
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) formResponse(r interface{}) ([]byte, error) {
	return json.Marshal(r)
}
