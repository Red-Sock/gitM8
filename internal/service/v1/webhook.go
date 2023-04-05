package v1

import (
	"gitM8/internal/service/domain/webhook"
)

type WebhookService struct {
}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (w *WebhookService) HandleWebhook(webhook webhook.Request) error {
	return nil
}
