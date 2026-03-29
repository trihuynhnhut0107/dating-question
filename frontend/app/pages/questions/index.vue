<!-- 4.5 app/pages/questions/index.vue - Infinite Scroll Page -->
<script setup lang="ts">
import { useQuestions } from "~/composables/useQuestions";
import type { QuestionDto } from "~/types/question";

const { fetchQuestionsByCursor } = useQuestions();
const questions = ref<QuestionDto[]>([]);
const nextCursor = ref<string | null>(null);
const hasNext = computed(() => !!nextCursor.value);

// Initial Load
const { data, pending } = await fetchQuestionsByCursor({ limit: 12 });
await console.log("Questions::", data.value);

// Initialize state when data arrives
watch(
  data,
  (newData) => {
    if (newData) {
      questions.value = newData.data;
      nextCursor.value = newData.next_cursor;
    }
  },
  { immediate: true },
);

// Method to fetch next page
const loadMore = async () => {
  if (!nextCursor.value || pending.value) return;

  // We use the same fetcher but with the next cursor
  const { data: nextPage } = await fetchQuestionsByCursor({
    cursor: nextCursor.value,
    limit: 12,
  });

  if (nextPage.value) {
    questions.value.push(...nextPage.value.data);
    nextCursor.value = nextPage.value.next_cursor;
  }
};
</script>

<template>
  <div class="questions-view">
    <h1>All Questions</h1>

    <div class="questions-grid">
      <QuestionsQuestionCard
        v-for="q in questions"
        :key="q.id"
        :content="q.content"
      />
    </div>

    <!-- Infinite Scroll Trigger -->
    <SharedLazyLoader
      :hasNext="hasNext"
      :pending="pending"
      @loadMore="loadMore"
    />
  </div>
</template>

<style scoped>
.questions-view {
  padding: 2rem;
  max-width: 1000px;
  margin: auto;
}
.questions-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(300px, 1fr));
  gap: 1.5rem;
  margin-top: 2rem;
}
</style>
