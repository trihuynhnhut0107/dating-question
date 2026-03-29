<!-- 3. app/components/shared/LazyLoader.vue - Infinite Scroll Loader -->
<script setup lang="ts">
const props = defineProps<{
  hasNext: boolean;
  pending: boolean;
}>();

const emit = defineEmits(["loadMore"]);

// Basic intersection observer to trigger loadMore
const loader = ref(null);
useIntersectionObserver(
  loader,
  ([entry]) => {
    if (entry?.isIntersecting && props.hasNext && !props.pending) {
      emit("loadMore");
    }
  },
  { rootMargin: "400px" }, // Trigger slightly before it enters the viewport
);
</script>

<template>
  <div ref="loader" class="lazy-loader">
    <div v-if="pending">Loading more...</div>
    <div v-else-if="!hasNext">No more questions</div>
  </div>
</template>

<style scoped>
.lazy-loader {
  padding: 2rem;
  text-align: center;
  color: #666;
}
</style>
