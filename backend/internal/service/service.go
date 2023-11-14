package service

import (
	"context"
	"webhook/internal/model"
)

type WebhookService interface {
	Create(ctx context.Context, webhook *model.Webhook) (*model.Webhook, error)
	Get(ctx context.Context, uuid string) (*model.Webhook, error)
	Delete(ctx context.Context, uuid string) error
}

type WebhookPublicService interface {
	GetWebhook(ctx context.Context, uuid string) (*model.Webhook, error)
}

type RequestService interface {
	List(ctx context.Context, webhookUuid string) ([]*model.Request, error)
	Dump(ctx context.Context, webhookUuid string, request *model.Request) (*model.Request, error)
}
