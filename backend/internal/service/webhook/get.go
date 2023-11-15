package webhook

import (
	"context"
	"errors"
	"webhook/internal/model"
	repoErrors "webhook/internal/repository/errors"
	"webhook/internal/repository/webhook/converter"
	serviceErrors "webhook/internal/service/errors"
)

func (s *service) Get(ctx context.Context, uuid string) (*model.Webhook, error) {
	webhook, err := s.webhookRepository.Get(ctx, uuid)
	if err != nil {
		if errors.Is(repoErrors.NotFoundError, err) {
			return nil, serviceErrors.NotFoundError
		}
		return nil, err
	}

	if webhook == nil {
		return nil, serviceErrors.NotFoundError
	}

	return converter.ToWebhookFromRepo(webhook), nil
}

func (s *service) GetWebhook(ctx context.Context, uuid string) (*model.Webhook, error) {
	return s.Get(ctx, uuid)
}
