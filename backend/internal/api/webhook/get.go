package webhook

import (
	"context"
	"webhook/internal/converter"
	desc "webhook/pkg"
)

func (i *Implementation) Get(ctx context.Context, uuid string) (*desc.GetWebhookResponse, error) {
	webhook, err := i.webhookService.Get(ctx, uuid)
	if err != nil {
		return nil, err
	}

	return &desc.GetWebhookResponse{
		Webhook: converter.ToWebhookFromService(webhook),
	}, nil
}
