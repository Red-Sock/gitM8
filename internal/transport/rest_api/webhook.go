package rest_api

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"

	ghmodel "github.com/Red-Sock/gitm8/internal/clients/git/model"
	"github.com/Red-Sock/gitm8/internal/service/domain"
)

var ErrParsingWebHookPath = errors.New("error parsing arguments from path on webhook request")

const (
	githubHeader = "X-GitHub-Event"
)

func (s *Server) Webhook(rw http.ResponseWriter, req *http.Request) {
	var ticket domain.TicketRequest
	var err error

	ticket.OwnerId, ticket.Uri, err = extractWebhookPath(req.URL.Path)
	if err != nil {
		logrus.Errorf("error extracting webhook path from url: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	payload, err := io.ReadAll(req.Body)
	if err != nil {
		logrus.Errorf("error reading webhook body: %s", err)
		rw.WriteHeader(http.StatusBadRequest)
		return
	}
	switch {
	case req.Header.Get(githubHeader) != "":
		var eventType domain.EventType
		eventType.ParseGithub(req.Header.Get(githubHeader))

		wh, err := ghmodel.SelectModel(eventType, payload)
		if err != nil {
			logrus.Errorf("error selecting proper webhook model: %s", err)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}

		ticket.Payload = wh
		ticket.RepoType = domain.RepoTypeGithub
	default:
		logrus.Errorf("error handling webhook: %s", fmt.Sprintf("no known webhook header is provided %v", req.Header))
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	err = s.services.WebhookService().HandleWebhook(ticket)
	if err != nil {
		logrus.Errorf("error handling webhook %s", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, _ = rw.Write([]byte(`OK`))
}

func extractWebhookPath(pth string) (ownerId uint64, ticketUUID string, err error) {
	pth = strings.TrimLeft(pth, "/")
	pathArgs := strings.Split(pth, "/")
	if len(pathArgs) < 2 {
		return 0, "", errors.Wrap(ErrParsingWebHookPath, "path should consist of 2 elements {{ user_id }} and {{ uuid_of_ticket }}")
	}

	ownerId, err = strconv.ParseUint(pathArgs[0], 10, 64)
	if err != nil {
		return 0, "", errors.Wrap(err, "error parsing ownerId from webhookRequest: %s")
	}

	ticketUUID = pathArgs[1]

	return ownerId, ticketUUID, nil
}
