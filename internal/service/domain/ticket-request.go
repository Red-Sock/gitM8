package domain

type TicketRequest struct {
	TicketId uint64
	OwnerId  uint64
	Uri      string
	RepoType RepoType
	Payload  Payload
}
