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
import { defineComponent, PropType, ref, computed } from 'vue';
import { Webhook } from "@/api";

export default defineComponent({
  props: {
    webhook: {
      type: Object as PropType<Webhook>,
      required: true,
    },
  },
  setup(props) {
    const snackbar = ref(false);
    const snackbarText = ref('');

    const webhookUrl = computed(() => {
      if (props.webhook && props.webhook.uuid) {
        return `http://localhost:8080/${props.webhook.uuid}`;
      }
      return ''; // Handle the case when webhook or uuid is missing
    });

    const copyUrlToClipboard = async () => {
      try {
        if (webhookUrl.value) {
          await navigator.clipboard.writeText(webhookUrl.value);
          snackbarText.value = 'URL copied to clipboard';
          snackbar.value = true;
        } else {
          snackbarText.value = 'Webhook data is missing';
          snackbar.value = true;
        }
      } catch (err) {
        snackbarText.value = 'Failed to copy URL';
        snackbar.value = true;
      }
    };

    return {
      webhookUrl,
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
  background-color: #ffffff;
}
</style>
