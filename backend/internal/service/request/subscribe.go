package request

import (
	"context"
	"webhook/internal/model"
	"webhook/internal/repository/request/converter"
)

func (s *service) Subscribe(ctx context.Context, uuid string) (<-chan *model.Request, error) {
	ch, _ := s.requestRepository.Subscribe(ctx, uuid)
	serviceCh := make(chan *model.Request)
	go func() {
		defer close(serviceCh)
		for {
			select {
			case <-ctx.Done():
				return
			case r := <-ch:
				serviceCh <- converter.ToRequestFromRepo(r)
			}
		}
	}()

	return serviceCh, nil
}
