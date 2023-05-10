package rest_api

import (
	"fmt"
	"io"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/service/domain/webhook"
)

const (
	githubHeader = "X-GitHub-Event"
)

func (s *Server) Webhook(_ http.ResponseWriter, req *http.Request) {
	var wh webhook.Request
	var err error

	wh.Payload, err = io.ReadAll(req.Body)
	if err != nil {
		logrus.Errorf("error reading webhook body: %s", err)
		return
	}

	switch {
	case req.Header.Get(githubHeader) != "":
		wh.Src = webhook.RepoTypeGithub
		wh.Type.ParseGithub(req.Header.Get(githubHeader))
	default:
		logrus.Errorf("error handling webhook: %s", fmt.Sprintf("no known webhook header is provided %v", req.Header))
		return
	}
	logrus.Infof("Payload: %s, Src: %d, Type: %d", string(wh.Payload), wh.Src, wh.Type)
	return

	err = s.services.WebhookService().HandleWebhook(wh)
	if err != nil {
		logrus.Errorf("error handling webhook %s", err.Error())
		return
	}
}
