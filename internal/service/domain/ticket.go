package domain

type Ticket struct {
	Id      uint64
	Name    string
	OwnerId uint64
	WebURL  string
}

type CreateTicketRequest struct {
	OwnerTgId uint64
}
