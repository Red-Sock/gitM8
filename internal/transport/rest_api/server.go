package rest_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"gitM8/internal/config"
	"gitM8/internal/service/interfaces"

	"github.com/gorilla/mux"
)

type Server struct {
	HttpServer *http.Server

	services interfaces.Services

	version string
}

func NewServer(cfg *config.Config, services interfaces.Services) *Server {
	r := mux.NewRouter()
	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + cfg.GetString(config.ServerRestAPIPort),
			Handler: r,
		},

		services: services,

		version: cfg.GetString(config.AppInfoVersion),
	}

	r.HandleFunc("/version", s.Version)
	r.HandleFunc("/webhooks", s.Webhook)
	return s
}

func (s *Server) Start(_ context.Context) error {
	go func() {
		err := s.HttpServer.ListenAndServe()
		if err != nil {
			log.Fatal(err)
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
