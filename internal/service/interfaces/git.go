package interfaces

import (
	"context"

	"gitM8/internal/service/domain"
)

type Git interface {
	GetCurrentUser(ctx context.Context) (domain.TgUser, error)
}
