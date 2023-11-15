package request

import (
	"context"
	"webhook/internal/model"
)

func (i *Implementation) Subscribe(ctx context.Context, uuid string) (<-chan *model.Request, error) {
	return i.requestService.Subscribe(ctx, uuid)
}
