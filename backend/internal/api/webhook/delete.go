package webhook

import (
	"context"
	desc "webhook/pkg"
)

func (i *Implementation) Delete(ctx context.Context, uuid string) (*desc.DeleteWebhookResponse, error) {
	err := i.webhookService.Delete(ctx, uuid)
	if err != nil {
		return nil, err
	}
	deleteResponse := &desc.DeleteWebhookResponse{
		Status: "ok",
	}

	return deleteResponse, nil
}
