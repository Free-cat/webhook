package webhook

import (
	"webhook/internal/service"
)

type Implementation struct {
	webhookService service.WebhookService
}

func NewImplementation(webhookService service.WebhookService) *Implementation {
	return &Implementation{
		webhookService: webhookService,
	}
}
