package create_ticket

import (
	"context"
	"fmt"

	tgapi "github.com/Red-Sock/go_tg/interfaces"
	"github.com/Red-Sock/go_tg/model"
	"github.com/Red-Sock/go_tg/model/response"

	"github.com/Red-Sock/gitM8/internal/service/domain"
	serviceInterfaces "github.com/Red-Sock/gitM8/internal/service/interfaces"
)

const Command = "/create-ticket"

type Handler struct {
	regService serviceInterfaces.RegistrationService
}

func New(regService serviceInterfaces.RegistrationService) *Handler {
	return &Handler{
		regService: regService,
	}
}

func (h *Handler) Handle(in *model.MessageIn, out tgapi.Chat) {
	resp, err := h.regService.CreateBasicTicket(context.Background(), domain.CreateTicketRequest{
		OwnerTgId: uint64(in.From.ID),
	})
	if err != nil {
		out.SendMessage(response.NewMessage("something went wrong: " + err.Error()))
		return
	}

	out.SendMessage(response.NewMessage(
		fmt.Sprintf("Insert this url as webhook for project: %s\nTicket id is: %d",
			resp.WebURL,
			resp.Id,
		)))
}
