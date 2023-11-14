package converter

import (
	"github.com/google/uuid"
	"time"
	"webhook/internal/model"
	desc "webhook/pkg"
)

func ToWebhookFromService(webhook *model.Webhook) *desc.Webhook {
	return &desc.Webhook{
		Uuid:      webhook.Uuid,
		CreatedAt: webhook.CreatedAt.Unix(),
	}
}

func ToWebhookFromDesc(webhook *desc.Webhook) *model.Webhook {
	webhookUuid := webhook.Uuid
	if webhookUuid == "" {
		webhookUuid = uuid.New().String()
	}
	return &model.Webhook{
		Uuid:      webhookUuid,
		CreatedAt: time.Unix(webhook.CreatedAt, 0),
	}
}
