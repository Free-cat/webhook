package webhook

import (
	"context"
	"github.com/google/uuid"
	"log"
	"webhook/internal/model"
	"webhook/internal/repository/webhook/converter"
)

func (s *service) Create(ctx context.Context, webhook *model.Webhook) (*model.Webhook, error) {
	webhookUuid := webhook.Uuid
	if webhookUuid == "" || webhook == nil {
		webhookUuid = uuid.New().String()
	}
	repoWebhook := converter.ToWebhookFromService(webhook)
	createdWebhook, err := s.webhookRepository.Create(ctx, webhookUuid, repoWebhook)
	if err != nil {
		log.Printf("Error while create Webhook: %v\n", err)
		return nil, nil
	}

	return converter.ToWebhookFromRepo(createdWebhook), nil
}
