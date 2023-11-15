import {ApiClient} from './client';
import {AxiosResponse} from "axios";

export async function fetchRequestsForWebhook(webhookUuid: string): Promise<Request[]> {
    try {
        const response: AxiosResponse<Request[]> = await ApiClient.get(`/requests/${webhookUuid}`);
        return response.data;
    } catch (error) {
        // Handle errors (e.g., network error, server error)
        throw error;
    }
}

// Function to establish a connection to a server-sent events (SSE) endpoint
export function fetchSSE(webhookUuid: string): EventSource {
    // Create a new EventSource instance pointing to the SSE endpoint
    const url = new URL(`/api/v1/requests/${webhookUuid}/sse`, ApiClient.defaults.baseURL);
    // The EventSource object is returned and can be used to listen to events
    return new EventSource(url.href);
}

