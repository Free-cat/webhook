export interface Webhook {
    uuid: string;
    created_at: number | null;
}

export interface Request {
    webhook_uuid: string;
    uuid: string;
    host: string;
    path: string;
    method: string;
    body: string;
    headers: Record<string, string[]>;
    created_at: number;
}

export interface CreateWebhookRequest {
    webhook: Webhook;
}

export interface CreateWebhookResponse {
    webhook: Webhook | null;
}

export interface GetWebhookResponse {
    webhook: Webhook | null;
}

export interface DeleteWebhookResponse {
    status: string;
}
