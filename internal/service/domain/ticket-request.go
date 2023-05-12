package domain

type TicketRequest struct {
	OwnerId uint64
	Uri     string
	Req     Request
}
