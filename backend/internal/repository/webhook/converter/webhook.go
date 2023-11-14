package converter

import (
	"webhook/internal/model"
	repoModel "webhook/internal/repository/webhook/model"
)

func ToWebhookFromRepo(webhook *repoModel.Webhook) *model.Webhook {
	return &model.Webhook{
		Uuid:      webhook.Uuid,
		CreatedAt: webhook.CreatedAt,
	}
}

func ToWebhookFromService(webhook *model.Webhook) *repoModel.Webhook {
	return &repoModel.Webhook{
		Uuid: webhook.Uuid,
	}
}
