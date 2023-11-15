<template>
  <v-container fluid>
    <v-row>
      <v-col cols="12" md="4" class="pa-4">
        <RequestsList
            :requests="requests"
            @request-selected="handleRequestSelected"
        />
      </v-col>
      <v-col cols="12" md="8" class="pa-4">
        <WebhookInfo :webhook="webhook" />
        <RequestDetails :request="selectedRequest" />
      </v-col>
    </v-row>
  </v-container>
</template>

<script lang="ts">
import {defineComponent, ref, onMounted, onUnmounted} from 'vue';
import { createWebhook, fetchWebhook } from '@/api/webhook';
import {fetchRequestsForWebhook} from '@/api/request';
import { useRoute } from 'vue-router';

import RequestsList from '@/components/RequestsList.vue';
import WebhookInfo from '@/components/WebhookInfo.vue';
import RequestDetails from '@/components/RequestDetails.vue';
import { Webhook, CreateWebhookRequest, Request } from "@/api";
import {ApiClient} from "@/api/client";

export default defineComponent({
  components: {
    RequestsList,
    WebhookInfo,
    RequestDetails
  },
  setup() {
    const requests = ref<Request[]>([]);
    const selectedRequest = ref(null);
    const route = useRoute();
    const webhook = ref<Webhook | null>(null);
    let eventSource: EventSource | null = null;

    function setupSSE(webhookUuid: string) {
      // Initialize the EventSource connection
      const url = new URL(`/api/v1/requests/sse/${webhookUuid}`, ApiClient.defaults.baseURL);
      eventSource = new EventSource(url.href);

      // Handle incoming messages
      eventSource.onmessage = (event) => {
        const request = JSON.parse(event.data);
        console.log('SSE message:', request)
        requests.value.unshift(request);
      };

      // Handle errors
      eventSource.onerror = (error) => {
        console.error('SSE error:', error);
        // Optionally, you could close the EventSource here
        // eventSource?.close();
      };

      eventSource.addEventListener('ping', (event) => {
        console.log('SSE ping:', event)
      });
    }

    async function fetchWebhookOrCreate(uuid: string) {
      try {
        // Try to fetch the existing webhook
        const existingWebhook = await fetchWebhook(uuid);
        webhook.value = existingWebhook.webhook;
      } catch (error) {
        console.error('Error fetching webhook:', error);

        try {
          const request: CreateWebhookRequest = {
            webhook: {
              uuid: uuid,
              created_at: null,
            },
          };
          const createdWebhook = await createWebhook(request);
          webhook.value = createdWebhook.webhook;
        } catch (createError) {
          console.error('Error creating webhook:', createError);

          webhook.value = null; // Set to null if creating also fails
        }
      }
    }

    function handleRequestSelected(request: any) {
      selectedRequest.value = request;
    }

    async function fetchRequests(webhookId: string) {
      try {
        const requestsForWebhook = await fetchRequestsForWebhook(webhookId);
        if (requestsForWebhook.length > 0) {
          requests.value = [];
        }
        requests.value = requestsForWebhook;
      } catch (error) {
        requests.value = [];
        console.error('Error fetching requests:', error);
      }
    }

    onMounted(async () => {
      const uuidFromRoute = route.params.uuid;
      await fetchWebhookOrCreate(uuidFromRoute);
      await fetchRequests(uuidFromRoute);
      setupSSE(uuidFromRoute);
    });

    onUnmounted(() => {
      console.log('Unmounting Webhook.vue')
      // Close the EventSource when the component unmounts to prevent memory leaks
      eventSource?.close();
    });

    return {
      requests,
      webhook,
      selectedRequest,
      handleRequestSelected
    };
  }
});
</script>

<style scoped>
.pa-4 {
  padding: 1rem !important;
}
</style>
