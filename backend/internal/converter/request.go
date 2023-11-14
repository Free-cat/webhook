package converter

import (
	"webhook/internal/model"
	desc "webhook/pkg"
)

func ToRequestFromService(request *model.Request) *desc.Request {
	return &desc.Request{
		WebhookUuid: request.WebhookUuid,
		Uuid:        request.Uuid,
		Path:        request.Path,
		Method:      request.Method,
		Body:        request.Body,
		Headers:     request.Headers,
		CreatedAt:   request.CreatedAt,
	}
}

func ToRequestFromDesc(request *desc.Request) *model.Request {
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

func ToRequestsFromService(requests []*model.Request) []*desc.Request {
	var result []*desc.Request
	for _, request := range requests {
		result = append(result, ToRequestFromService(request))
	}
	return result
}
