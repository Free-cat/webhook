package webhook

import (
	"context"
	"webhook/internal/converter"
	desc "webhook/pkg"
)

func (i *Implementation) Create(ctx context.Context, req *desc.CreateWebhookRequest) (*desc.CreateWebhookResponse, error) {
	webhook, err := i.webhookService.Create(ctx, converter.ToWebhookFromDesc(&req.Webhook))
	if err != nil {
		return nil, err
	}

	return &desc.CreateWebhookResponse{
		Webhook: converter.ToWebhookFromService(webhook),
	}, nil
}
