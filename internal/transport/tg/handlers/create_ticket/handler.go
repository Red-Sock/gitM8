package create_ticket

import (
	"context"
	"fmt"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"
	"github.com/sirupsen/logrus"

	"github.com/Red-Sock/gitm8/internal/service/domain"
	serviceInterfaces "github.com/Red-Sock/gitm8/internal/service/interfaces"
)

const Command = "/create_ticket"

type Handler struct {
	regService serviceInterfaces.TicketsService
}

func New(regService serviceInterfaces.TicketsService) *Handler {
	return &Handler{
		regService: regService,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	resp, err := h.regService.CreateBasicTicket(context.Background(), domain.CreateTicketRequest{
		OwnerTgId: uint64(in.From.ID),
	})
	if err != nil {
		logrus.Errorf("error creating basic ticket: %s", err)
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}

	webUrl, err := resp.GetWebUrl()
	if err != nil {
		logrus.Errorf("error getting weburl of ticket: %s", err)
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}

	out.SendMessage(response.NewMessage(fmt.Sprintf("Insert this url as webhook for project: %s\nTicket id is: %d",
		webUrl,
		resp.Id)))
}

func (h *Handler) GetDescription() string {
	return "Create ticket"
}
