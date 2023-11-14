package webhook

import (
	"context"
	"webhook/internal/model"
	"webhook/internal/repository/webhook/converter"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.Webhook, error) {
	webhook, err := s.webhookRepository.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}
	if webhook == nil {
		return nil, nil
	}

	return converter.ToWebhookFromRepo(webhook), nil
}

func (s *service) GetWebhook(ctx context.Context, uuid string) (*model.Webhook, error) {
	return s.Get(ctx, uuid)
}
