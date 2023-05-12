package rest_api

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/service/domain"
)

const (
	githubHeader = "X-GitHub-Event"
)

func (s *Server) Webhook(rw http.ResponseWriter, req *http.Request) {
	var ticket domain.TicketRequest
	var err error

	ticket.OwnerId, ticket.Uri, err = extractWebhookPath(req.URL.Path)
	if err != nil {
		logrus.Error(err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	ticket.Req.Payload, err = io.ReadAll(req.Body)
	if err != nil {
		logrus.Errorf("error reading webhook body: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	switch {
	case req.Header.Get(githubHeader) != "":
		ticket.Req.Src = domain.RepoTypeGithub
		ticket.Req.Type.ParseGithub(req.Header.Get(githubHeader))
	default:
		logrus.Errorf("error handling webhook: %s", fmt.Sprintf("no known webhook header is provided %v", req.Header))
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	logrus.Infof("Payload: %s, Src: %d, Type: %d", string(ticket.Req.Payload), ticket.Req.Src, ticket.Req.Type)

	err = s.services.WebhookService().HandleWebhook(ticket)
	if err != nil {
		logrus.Errorf("error handling webhook %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func extractWebhookPath(pth string) (ownerId uint64, ticketUUID string, err error) {
	pathArgs := strings.Split(pth, "/")
	if len(pathArgs) < 4 {
		return 0, "", errors.New("error parsing arguments from path on webhook request")
	}

	pathArgs = pathArgs[2:]

	ownerId, err = strconv.ParseUint(pathArgs[0], 10, 10)
	if err != nil {
		return 0, "", errors.Join(errors.New("error parsing ownerId from webhookRequest: %s"), err)
	}

	ticketUUID = pathArgs[1]

	return ownerId, ticketUUID, nil
}
