package pkg

type Request struct {
	WebhookUuid string              `json:"webhook_uuid"`
	Uuid        string              `json:"uuid"`
	Host        string              `json:"host"`
	Path        string              `json:"path"`
	Method      string              `json:"method"`
	Body        string              `json:"body"`
	Headers     map[string][]string `json:"headers"`
	CreatedAt   int64               `json:"created_at"`
}

type DumpResponse struct {
	Request *Request `json:"request"`
}

type ErrWebhookNotFound struct{}

func (e *ErrWebhookNotFound) Error() string {
	return "webhook not found"
}
