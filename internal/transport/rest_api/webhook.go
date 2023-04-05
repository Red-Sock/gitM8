package rest_api

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"google.golang.org/appengine/log"

	"gitM8/internal/service/domain/webhook"
)

const (
	githubHeader = "X-GitHub-Event"
)

func (s *Server) Webhook(_ http.ResponseWriter, req *http.Request) {
	ctx := context.Background()

	var wh webhook.Request
	var err error

	wh.Payload, err = io.ReadAll(req.Body)
	if err != nil {
		log.Errorf(ctx, "error reading webhook body: %s", err)
		return
	}

	switch {
	case req.Header.Get(githubHeader) != "":
		wh.Src = webhook.RepoTypeGithub
		wh.Type.ParseGithub(req.Header.Get(githubHeader))
	default:
		log.Errorf(ctx, "error handling webhook: %s", fmt.Sprintf("no known webhook header is provided %v", req.Header))
		return
	}
	log.Infof(ctx, "Payload: %s, Src: %s, Type: %s", string(wh.Payload), wh.Src, wh.Type)
	return

	err = s.services.WebhookService().HandleWebhook(wh)
	if err != nil {
		log.Errorf(ctx, "error handling webhook %s", err.Error())
		return
	}
}
