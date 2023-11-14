<template>
  <v-card elevation="2" class="mb-4 info-card">
    <v-card-title class="headline">Your unique URL</v-card-title>
    <v-card-text>
      <div class="d-flex align-center">
        <v-tooltip bottom>
          <template v-slot:activator="{ on, attrs }">
            <code v-bind="attrs" v-on="on" @click="copyUrlToClipboard" class="mr-2">{{ webhookUrl }}</code>
          </template>
          <span>Copy to clipboard</span>
        </v-tooltip>
        <v-btn icon x-small @click="copyUrlToClipboard" class="mr-2" size="x-small">
          <v-icon size="small">mdi-content-copy</v-icon>
        </v-btn>
        <v-btn icon :href="webhookUrl" target="_blank" class="mr-2" size="x-small">
          <v-icon size="small">mdi-open-in-new</v-icon>
        </v-btn>
      </div>
    </v-card-text>
  </v-card>
  <v-snackbar v-model="snackbar" :timeout="3000" bottom right>
    {{ snackbarText }}
    <template v-slot:action="{ attrs }">
      <v-btn color="blue" text v-bind="attrs" @click="snackbar = false">Close</v-btn>
    </template>
  </v-snackbar>
</template>

<script lang="ts">
import { defineComponent, ref } from 'vue';

export default defineComponent({
  setup() {
    const hasRequests = ref(false); // Assume no requests to start with
    const webhookUrl = ref('https://webhook.site/unique-id');
    const webhookEmail = ref('unique-id@email.webhook.site');
    const snackbar = ref(false);
    const snackbarText = ref('');

    const copyUrlToClipboard = async () => {
      try {
        await navigator.clipboard.writeText(webhookUrl.value);
        snackbarText.value = 'URL copied to clipboard';
        snackbar.value = true;
      } catch (err) {
        snackbarText.value = 'Failed to copy URL';
        snackbar.value = true;
      }
    };

    return {
      hasRequests,
      webhookUrl,
      webhookEmail,
      snackbar,
      snackbarText,
      copyUrlToClipboard
    };
  }
});
</script>

<style scoped>
code {
  background-color: #f0f0f0;
  padding: 4px 6px;
  border-radius: 4px;
  cursor: pointer;
  user-select: none;
}
.headline {
  background-color: #e8e8e8;
  color: #424242;
  margin-bottom: 10px;
}
.info-card {
  background-color: #f5f5f5;
}
</style>
