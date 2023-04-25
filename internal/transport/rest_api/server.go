package rest_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/config"
	"github.com/Red-Sock/gitm8/internal/service/interfaces"

	"github.com/gorilla/mux"
)

type Server struct {
	HttpServer *http.Server

	services interfaces.Services

	version string
}

func NewServer(cfg *config.Config, services interfaces.Services) (*Server, error) {
	port, err := cfg.TryGetInt(config.ServerRestAPIPort)
	if err != nil {
		return nil, errors.Wrap(err, "error extracting "+config.ServerRestAPIPort+" from config")
	}
	r := mux.NewRouter()
	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + strconv.Itoa(port),
			Handler: r,
		},

		services: services,

		version: cfg.GetString(config.AppInfoVersion),
	}

	r.HandleFunc("/version", s.Version)
	r.HandleFunc("/webhooks", s.Webhook)
	return s, nil
}

func (s *Server) Start(_ context.Context) error {
	go func() {
		err := s.HttpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	logrus.Infof("started http listener on: %s", s.HttpServer.Addr)

	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.HttpServer.Shutdown(ctx)
}

func (s *Server) formResponse(r interface{}) ([]byte, error) {
	return json.Marshal(r)
}
