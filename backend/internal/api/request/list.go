package request

import (
	"context"
	"webhook/internal/converter"
	desc "webhook/pkg"
)

func (i *Implementation) List(ctx context.Context, webhookUuid string) ([]*desc.Request, error) {
	requests, err := i.requestService.List(ctx, webhookUuid)
	if err != nil {
		return nil, err
	}

	return converter.ToRequestsFromService(requests), nil
}
