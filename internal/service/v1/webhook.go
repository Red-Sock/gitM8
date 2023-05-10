package v1

import (
	"github.com/Red-Sock/gitm8/internal/service/domain/webhook"
)

type WebhookService struct {
}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (w *WebhookService) HandleWebhook(webhook webhook.Request) error {
	return nil
}
