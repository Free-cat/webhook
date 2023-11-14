<template>
  <v-card elevation="2" class="request-list">
    <v-card-title class="headline">Requests</v-card-title>
    <v-card-text>
      <v-list dense>
        <v-list-item-group v-model="selectedRequestIndex" color="blue lighten-3">
          <v-list-item
              v-for="(request, index) in requests"
              :key="request.id"
              @click="selectRequest(request)"
          >
            <v-list-item-content>
              <v-list-item-title><v-chip label>{{ request.method }}</v-chip> {{ request.path }}</v-list-item-title>
              <v-list-item-subtitle>{{ request.timestamp }}</v-list-item-subtitle>
            </v-list-item-content>
          </v-list-item>
        </v-list-item-group>
      </v-list>
    </v-card-text>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from 'vue';

export default defineComponent({
    props: {
        requests: {
            type: Array as PropType<Array<any>>,
            required: true
        }
    },
    setup(props, { emit }) {
        const selectedRequestIndex = ref<number | null>(null);

        function selectRequest(request: any) {
            emit('request-selected', request);
        }

        return {
            selectedRequestIndex,
            selectRequest
        };
    }
});
</script>

<style scoped>
.request-list {
  max-height: 600px;
  overflow-y: auto;
}
.headline {
  background-color: #e8e8e8;
  color: #424242;
}
</style>
