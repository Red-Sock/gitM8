package create_ticket

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
	"github.com/Red-Sock/gitm8/internal/transport/tg/commands"
	"github.com/Red-Sock/gitm8/internal/transport/tg/constructors"
)

type Handler struct {
	ticketServices serviceInterfaces.TicketsService
	host           string
}

func (h *Handler) GetCommand() string {
	return commands.CreateTicket
}

func New(srv serviceInterfaces.Services, host string) *Handler {
	return &Handler{
		ticketServices: srv.TicketsService(),
		host:           host,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	resp, err := h.ticketServices.CreateBasicTicket(context.Background(), domain.CreateTicketRequest{
		OwnerTgId: uint64(in.From.ID),
	})
	if err != nil {
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}

	webUrl, err := resp.GetWebUrl()
	if err != nil {
		logrus.Errorf("error getting weburl of ticket: %s", err)
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}
	out.SendMessage(&response.DeleteMessage{
		ChatId:    in.Chat.ID,
		MessageId: int64(in.MessageID),
	})

	webUrl, err = url.JoinPath(h.host, webUrl)
	if err != nil {
		logrus.Errorf("error creating weburl of ticket: %s", err)
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}

	out.SendMessage(constructors.GetEndState(fmt.Sprintf("Insert this url as webhook for project: " + webUrl + "\nTicket id is: " + strconv.FormatUint(resp.Id, 10))))
}

func (h *Handler) GetDescription() string {
	return "Create ticket"
}
