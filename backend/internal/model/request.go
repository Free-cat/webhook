package model

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
