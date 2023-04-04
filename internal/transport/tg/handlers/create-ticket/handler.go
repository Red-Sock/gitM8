package create_ticket

import (
	"context"
	"fmt"

	tgapi "github.com/AlexSkilled/go_tg/interfaces"
	"github.com/AlexSkilled/go_tg/model"
	"github.com/AlexSkilled/go_tg/model/response"

	"gitM8/internal/service/domain"
	serviceInterfaces "gitM8/internal/service/interfaces"
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
		OwnerId: uint64(in.Contact.UserID),
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
