package repository

import (
	"context"
	repoRequestModel "webhook/internal/repository/request/model"
	repoWebhookModel "webhook/internal/repository/webhook/model"
)

type RequestRepository interface {
	List(ctx context.Context, webhookUuid string) ([]*repoRequestModel.Request, error)
	Create(ctx context.Context, webhookUuid string, request *repoRequestModel.Request) (*repoRequestModel.Request, error)
	Get(ctx context.Context, webhookUuid string, uuid string) (*repoRequestModel.Request, error)
	Delete(ctx context.Context, webhookUuid string, uuid string) error
}

type WebhookRepository interface {
	Create(ctx context.Context, uuid string, webhook *repoWebhookModel.Webhook) (*repoWebhookModel.Webhook, error)
	Get(ctx context.Context, uuid string) (*repoWebhookModel.Webhook, error)
	Delete(ctx context.Context, uuid string) error
}
