package request

import (
	"context"
	"webhook/internal/model"
	"webhook/internal/repository/request/converter"
)

func (s *service) List(ctx context.Context, webhookUuid string) ([]*model.Request, error) {
	requests, err := s.requestRepository.List(ctx, webhookUuid)
	if err != nil {
		return nil, err
	}

	return converter.ToRequestsFromRepo(requests), nil
}
