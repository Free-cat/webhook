<template>
  <v-card elevation="2">
    <v-card-title class="headline d-flex justify-space-between align-center">
      <span>Requests</span>
      <v-icon small color="green">mdi-check-circle</v-icon>
    </v-card-title>
    <v-list class="request-list">
      <v-list-item-group v-model="selectedRequestIndex" color="blue lighten-3">
        <v-list-item
            v-for="(request, index) in requests"
            :key="request.uuid"
            @click="selectRequest(request)"
            :class="index === selectedRequestIndex ? 'selected' : ''"
        >
          <v-list-item-content>
            <v-list-item-title ><v-chip label>{{ request.method }}</v-chip> #{{ request.uuid.slice(-5) }}</v-list-item-title>
            <v-list-item-subtitle >{{ new Date(request.created_at * 1000).toLocaleString() }}</v-list-item-subtitle>
          </v-list-item-content>
        </v-list-item>
      </v-list-item-group>
    </v-list>
  </v-card>
</template>

<script lang="ts">
import { defineComponent, PropType, ref } from 'vue';
import { Request } from '@/api'; // Import the listRequests function

export default defineComponent({
  props: {
    requests: {
      type: Array as PropType<Request[]>,
      required: true
    },
    eventSource: null,
  },
  setup(props, { emit }) {
    const selectedRequestIndex = ref<number | null>(null);

    function selectRequest(request: Request) {
      selectedRequestIndex.value = props.requests.findIndex((r) => r.uuid === request.uuid);
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
  max-height: 80vh;
  font-size: 90%!important;
}
.headline {
  background-color: #e8e8e8;
  color: #424242;
}
.selected{
  background-color: #eeeeee;
  color: #424242;
}
</style>
