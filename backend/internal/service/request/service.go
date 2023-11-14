package request

import (
	"webhook/internal/repository"
	def "webhook/internal/service"
)

var _ def.RequestService = (*service)(nil)

type service struct {
	requestRepository repository.RequestRepository
}

func NewService(
	requestRepository repository.RequestRepository,
) *service {
	return &service{
		requestRepository: requestRepository,
	}
}
