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
import { defineComponent, ref } from 'vue';
import RequestsList from '@/components/RequestsList.vue';
import WebhookInfo from '@/components/WebhookInfo.vue';
import RequestDetails from '@/components/RequestDetails.vue';

export default defineComponent({
  components: {
    RequestsList,
    WebhookInfo,
    RequestDetails
  },
  setup() {
    const requests = ref([
      { id: 1, method: 'GET', path: '/api/data', timestamp: '2023-11-14T12:34:56Z' },
      { id: 2, method: 'POST', path: '/api/submit', timestamp: '2023-11-14T12:35:10Z' },
    ]);
    const webhook = ref({
      id: '123456',
      endpoint: 'http://example.com/webhook',
    });
    const selectedRequest = ref(null);

    function handleRequestSelected(request: any) {
      selectedRequest.value = request;
    }

    function fetchRequests() {
      // Fetching logic here
    }

    fetchRequests();

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
