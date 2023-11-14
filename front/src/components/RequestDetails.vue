<template>
  <v-card elevation="2" class="details-card">
    <v-card-title class="headline">Request Details</v-card-title>
    <v-card-text v-if="request">
      <div><strong>Method:</strong> {{ request.method }}</div>
      <div><strong>URL:</strong> {{ request.url }}</div>
      <div><strong>Status:</strong> {{ request.status }}</div>
      <div><strong>Headers:</strong> <pre>{{ request.headers }}</pre></div>
      <div><strong>Body:</strong> <pre>{{ request.body }}</pre></div>
      <!-- More details here -->
    </v-card-text>
    <v-card-text v-else>No request selected</v-card-text>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';

export default defineComponent({
    props: {
        request: {
            type: Object as PropType<any>,
            default: null
        }
    }
});
</script>

<style scoped>
.details-card {
  margin-top: 1rem;
}
.headline {
  background-color: #1976d2;
  color: white;
  margin-bottom: 10px;
}
pre {
  background-color: #eceff1;
  padding: 10px;
  border-radius: 4px;
}
</style>
