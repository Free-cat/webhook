package pkg

type Webhook struct {
	Uuid      string `json:"uuid"`
	CreatedAt int64  `json:"created_at"`
}

type CreateWebhookRequest struct {
	Webhook Webhook `json:"webhook"`
}

type CreateWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

type GetWebhookResponse struct {
	Webhook *Webhook `json:"webhook"`
}

type DeleteWebhookResponse struct {
	Status string `json:"status"`
}
