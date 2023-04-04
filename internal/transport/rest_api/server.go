package rest_api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"gitM8/internal/config"

	"github.com/gorilla/mux"
)

type Server struct {
	HttpServer *http.Server

	version string
}

func NewServer(cfg *config.Config) *Server {
	r := mux.NewRouter()
	s := &Server{
		HttpServer: &http.Server{
			Addr:    "0.0.0.0:" + cfg.GetString(config.ServerRestAPIPort),
			Handler: r,
		},

		version: cfg.GetString(config.AppInfoVersion),
	}

	r.HandleFunc("/version", s.Version)

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
