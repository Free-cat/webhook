<template>
  <div v-if="request">
    <v-row>
      <v-col cols="12" md="12">
        <v-card elevation="2" class="details-card">
          <v-card-title class="headline">
            Request Details
          </v-card-title>
          <v-card-text>
            <v-table density="compact">
              <template v-slot:bottom>
                <tbody>
                <tr>
                  <td class="pr-sm-10"><v-chip label style="font-size: 90%" size="x-small">{{request.method}}</v-chip></td>
                  <td><a href="https://webhook.site/bc7e5b7d-f136-470f-8fe3-9190741f33d2"> https://webhook.site/bc7e5b7d-f136-470f-8fe3-9190741f33d2</a></td>
                </tr>
                <tr>
                  <td class="pr-sm-10"><strong>UUID:</strong></td>
                  <td>{{request.uuid}}</td>
                </tr>
                <tr>
                  <td class="pr-sm-10"><strong>Date:</strong></td>
                  <td>{{ new Date(request.created_at * 1000).toLocaleString() }}</td>
                </tr>
                </tbody>
              </template>
            </v-table>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="12">
        <v-card elevation="2" class="details-card">
          <v-card-title class="headline">
            Headers
          </v-card-title>
          <v-card-text>
            <v-table density="compact">
              <template v-slot:bottom>
                <tbody>
                <tr v-for="(value, name) in request.headers" :key="name">
                  <td class="pr-sm-10">{{ name }}</td>
                  <td>{{ value[0] }}</td>
                </tr>
                </tbody>
              </template>
            </v-table>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col cols="12" md="12">
        <v-card elevation="2" class="details-card">
          <v-card-title class="headline">
            Body
          </v-card-title>
          <v-card-text>
            <v-code>
              <pre>{{ bodyToJson(request.body) }}</pre>
            </v-code>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </div>
</template>

<script lang="ts">
import { defineComponent, PropType } from 'vue';
import { Request } from '@/api';

export default defineComponent({
  props: {
    request: {
      type: Object as PropType<Request>,
      default: null
    }
  },
  methods:{
    bodyToJson(body: string) {
      try {
        return JSON.stringify(JSON.parse(body), null, 4);
      } catch (error) {
        return body;
      }
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
.details-card{
  font-size: 80%!important;
}
</style>
