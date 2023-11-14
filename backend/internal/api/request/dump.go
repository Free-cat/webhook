package request

import (
	"context"
	"log"
	"webhook/internal/converter"
	desc "webhook/pkg"
)

func (i *Implementation) Dump(ctx context.Context, webhookUuid string, request *desc.Request) (*desc.DumpResponse, error) {
	webhook, err := i.webhookPublicService.GetWebhook(ctx, webhookUuid)
	if webhook == nil || err != nil {
		return nil, &desc.ErrWebhookNotFound{}
	}

	request.WebhookUuid = webhookUuid
	dump, err := i.requestService.Dump(ctx, webhookUuid, converter.ToRequestFromDesc(request))
	if err != nil {
		log.Printf("requestService.Dump: %v", err)
		return nil, err
	}

	return &desc.DumpResponse{
		Request: converter.ToRequestFromService(dump),
	}, nil
}
