package interfaces

import (
	tgapi "github.com/Red-Sock/go_tg/interfaces"
)

type Chat interface {
	Send(out tgapi.MessageOut)
}
