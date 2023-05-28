package domain

import (
	"net/url"
	"strconv"
)

type Ticket struct {
	Id        uint64
	Name      string
	OwnerId   uint64
	URI       string
	GitSystem RepoType
}

func (t *Ticket) GetWebUrl() (string, error) {
	return url.JoinPath(strconv.Itoa(int(t.OwnerId)), t.URI)
}

type CreateTicketRequest struct {
	OwnerTgId uint64
	ChatId    uint64
}
