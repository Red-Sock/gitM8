package interfaces

import (
	"context"

	"github.com/Red-Sock/gitM8/internal/service/domain"
)

type Git interface {
	GetCurrentUser(ctx context.Context) (domain.TgUser, error)
}
