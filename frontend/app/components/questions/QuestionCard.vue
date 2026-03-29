<!-- 5. app/components/questions/QuestionCard.vue - Visual Card -->
<template>
  <div class="card">
    <p class="content">{{ content }}</p>
  </div>
</template>

<script setup lang="ts">
defineProps<{ content: string }>();
</script>

<style scoped>
.card {
  background: white;
  padding: 3rem;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
  min-height: 200px;
  display: flex;
  align-items: center;
  justify-content: center;
  text-align: center;
  font-size: 1.5rem;
  font-weight: 500;
  color: #2f3542;
}
</style>
