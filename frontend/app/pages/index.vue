<script setup lang="ts">
const loading = ref(false);
const { createSession } = useSessions();

async function startSession() {
  loading.value = true;

  try {
    const { data } = await createSession({ name: "Something" });
    if (data.value) {
      console.log(data.value);
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
      <NuxtLink to="/questions" class="btn btn-primary">
        Browse All Questions
      </NuxtLink>
      <button class="btn btn-primary" @click="startSession" :disabled="loading">
        {{ loading ? "Creating..." : "Random Session" }}
      </button>
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
.actions {
  display: flex;
  gap: 1.25rem;
  margin-top: 1.5rem;
  width: 100%;
  max-width: 500px;
  justify-content: center;
}

.btn {
  display: inline-flex;
  align-items: center;
  justify-content: center;
  padding: 0 2rem;
  height: 52px;
  border-radius: 12px; /* Slightly more rounded for a modern look */
  text-decoration: none;
  font-family: inherit;
  font-size: 1rem;
  font-weight: 600;
  cursor: pointer;
  border: none;
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
  box-sizing: border-box;
  flex: 1; /* Both buttons grow equally by default */
  min-width: 180px;
}

.btn:active {
  transform: translateY(1px) scale(0.98);
}

.btn-primary {
  background: #ff4757;
  color: white;
  box-shadow: 0 4px 12px rgba(255, 71, 87, 0.3);
}

.btn-primary:hover:not(:disabled) {
  background: #ff6b81;
  box-shadow: 0 6px 16px rgba(255, 71, 87, 0.4);
}

.btn-secondary {
  background: #2f3542;
  color: white;
}

.btn-secondary:hover:not(:disabled) {
  background: #57606f;
}

.btn:disabled {
  opacity: 0.6;
  cursor: not-allowed;
}

/* Responsive Styles */
@media (max-width: 600px) {
  .actions {
    flex-direction: column;
    padding: 0 1.5rem;
    gap: 1rem;
  }

  .btn {
    width: 100%;
    height: 48px; /* Slightly shorter on mobile */
    min-width: unset;
  }
}
</style>
