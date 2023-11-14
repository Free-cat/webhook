package converter

import (
	"webhook/internal/model"
	repoModel "webhook/internal/repository/request/model"
)

func ToRequestFromService(request *model.Request) *repoModel.Request {
	return &repoModel.Request{
		WebhookUuid: request.WebhookUuid,
		Uuid:        request.Uuid,
		Path:        request.Path,
		Method:      request.Method,
		Body:        request.Body,
		Headers:     request.Headers,
		CreatedAt:   request.CreatedAt,
	}
}

func ToRequestFromRepo(request *repoModel.Request) *model.Request {
	return &model.Request{
		WebhookUuid: request.WebhookUuid,
		Uuid:        request.Uuid,
		Path:        request.Path,
		Method:      request.Method,
		Body:        request.Body,
		Headers:     request.Headers,
		CreatedAt:   request.CreatedAt,
	}
}

func ToRequestsFromRepo(requests []*repoModel.Request) []*model.Request {
	var result []*model.Request
	for _, request := range requests {
		result = append(result, ToRequestFromRepo(request))
	}
	return result
}
