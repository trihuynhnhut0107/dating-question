<script setup lang="ts">
const route = useRoute();
const sessionId = route.params.id;
</script>

<template>
  <div>
    <p>Session ID: {{ sessionId }}</p>
  </div>
</template>
