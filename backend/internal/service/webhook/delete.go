package webhook

import "context"

func (s *service) Delete(ctx context.Context, uuid string) error {
	return s.webhookRepository.Delete(ctx, uuid)
}
