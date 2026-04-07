<script setup lang="ts">
const route = useRoute();
const sessionId = route.params.id;

const { fetchQuestionsByCursor } = useQuestions();

const {
  data: questions,
  limit,
  next_cursor,
} = await fetchQuestionsByCursor({
  session_id: String(sessionId),
});
</script>

<template>
  <div>
    <ul>
      <li v-for="question in (questions as any)" :key="question.id">
        {{ (question as any).content }}
      </li>
    </ul>
  </div>
</template>
