package request

import "webhook/internal/service"

type Implementation struct {
	requestService       service.RequestService
	webhookPublicService service.WebhookPublicService
}

func NewImplementation(requestService service.RequestService, webhookPublicService service.WebhookPublicService) *Implementation {
	return &Implementation{
		requestService:       requestService,
		webhookPublicService: webhookPublicService,
	}
}
