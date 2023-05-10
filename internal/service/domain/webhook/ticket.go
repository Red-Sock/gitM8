package webhook

type Ticket struct {
	OwnerId   int
	Timestamp string
	Req       Request
}
