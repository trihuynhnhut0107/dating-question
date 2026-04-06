<script setup lang="ts">
const loading = ref(false);
const { createSession } = useSessions();

async function startSession() {
  loading.value = true;

  try {
    const { data, error } = await createSession({ name: "Something" });
    if (data.value) {
      await navigateTo(`/session/${data.value.id}`);
    }
  } catch (err) {
    console.error("Failed to create session:", err);
  } finally {
    loading.value = false;
  }
}
</script>

<template>
  <div class="home-container">
    <h1>Dating Questions</h1>
    <p>Break the ice with meaningful conversations.</p>
    <div class="actions">
      <NuxtLink to="/questions" class="btn btn-primary"
        >Browse All Questions</NuxtLink
      >
      <button class="btn btn-create-session">Random Session</button>
    </div>
  </div>
</template>

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

.btn-create-session {
  background: #2f3542;
  color: white;
}
</style>
