package request

import (
	"context"
	"webhook/internal/model"
	repoConverter "webhook/internal/repository/request/converter"
)

func (s *service) Dump(ctx context.Context, webhookUuid string, request *model.Request) (*model.Request, error) {
	repoRequest, err := s.requestRepository.Create(ctx, webhookUuid, repoConverter.ToRequestFromService(request))
	if err != nil {
		return nil, err
	}

	return repoConverter.ToRequestFromRepo(repoRequest), nil
}
