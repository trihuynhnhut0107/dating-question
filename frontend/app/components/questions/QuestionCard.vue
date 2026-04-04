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
  --border-w: 2px;
  position: relative;
  background: #3a506b;
  padding-left: 3rem;
  padding-right: 3rem;
  padding-top: 1rem;
  padding-bottom: 1rem;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.1);
  margin-bottom: 2rem;
  display: flex;
  align-items: center;
  justify-content: start;
  text-align: start;
  font-size: 1.5rem;
  font-weight: 500;
  color: white;
  border: var(--border-w) solid #5bc0be;
  overflow: hidden;
}
.card:hover {
  cursor: pointer;
}
.card::after {
  content: "";
  position: absolute;

  left: 50%;
  bottom: 0;

  width: calc(100% - 4px);
  height: 2px;

  background-color: #6fffe9;

  transform: translateX(-50%) scaleX(0);
  transform-origin: center;
  transition: transform 0.3s ease;
}
.card:hover::after {
  transform: translateX(-50%) scaleX(1);
}
</style>
