<template>
</template>

<script lang="ts">
import { defineComponent, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import { v4 as uuidv4 } from 'uuid';

export default defineComponent({
  setup() {
    const router = useRouter();

    onMounted(() => {
      let uuid = localStorage.getItem('uuid');

      if (!uuid) {
        // If UUID doesn't exist in localStorage, generate a new one
        uuid = uuidv4();
        localStorage.setItem('uuid', uuid);
      }

      router.push({ name: 'Webhook', params: { uuid } });
    });

    return {};
  }
});
</script>
