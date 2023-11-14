package webhook

import (
	"webhook/internal/repository"
	def "webhook/internal/service"
)

var _ def.WebhookService = (*service)(nil)
var _ def.WebhookPublicService = (*service)(nil)

type service struct {
	webhookRepository repository.WebhookRepository
}

func NewService(
	WebhookRepository repository.WebhookRepository,
) *service {
	return &service{
		webhookRepository: WebhookRepository,
	}
}

func NewPublicService(
	WebhookRepository repository.WebhookRepository,
) *service {
	return &service{
		webhookRepository: WebhookRepository,
	}
}
