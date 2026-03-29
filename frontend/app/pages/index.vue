<template>
  <div class="home-container">
    <h1>Dating Questions</h1>
    <p>Break the ice with meaningful conversations.</p>
    <div class="actions">
      <NuxtLink to="/questions" class="btn btn-primary"
        >Browse All Questions</NuxtLink
      >
      <!-- <NuxtLink to="/session" class="btn btn-secondary"
        >Random Session</NuxtLink
      > -->
    </div>
  </div>
</template>

<script setup lang="ts"></script>

<style scoped>
.home-container {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  height: 80vh;
  text-align: center;
}
.btn {
  padding: 1rem 2rem;
  border-radius: 8px;
  text-decoration: none;
  margin: 0.5rem;
}
.btn-primary {
  background: #ff4757;
  color: white;
}
.btn-secondary {
  background: #2f3542;
  color: white;
}
</style>
