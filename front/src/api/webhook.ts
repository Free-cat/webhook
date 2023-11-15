// api/webhook.ts
import {ApiClient} from "@/api/client";
import {CreateWebhookRequest, CreateWebhookResponse, DeleteWebhookResponse, GetWebhookResponse} from "@/api";
import {AxiosResponse} from "axios";

export const createWebhook = async (request: CreateWebhookRequest): Promise<CreateWebhookResponse> => {
    try {
        const response:AxiosResponse<CreateWebhookResponse> = await ApiClient.post('/webhooks/', request);
        return response.data;
    } catch (error) {
        // Handle error
        throw error;
    }
};

export const fetchWebhook = async (uuid: string): Promise<GetWebhookResponse> => {
    try{
        const response:AxiosResponse<GetWebhookResponse> = await ApiClient.get(`/webhooks/${uuid}`);
        return response.data;
    } catch (error) {
        // Handle error
        throw error;
    }
};

export const deleteWebhook = async (uuid: string): Promise<DeleteWebhookResponse> => {
    try {
        const response:AxiosResponse<DeleteWebhookResponse> = await ApiClient.delete(`/webhooks/${uuid}`);
        return response.data;
    } catch (error) {
        // Handle error
        throw error;
    }
};
