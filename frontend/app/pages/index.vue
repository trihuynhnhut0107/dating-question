<script setup lang="ts">
definePageMeta({ title: 'Home' })

const features = [
  { icon: 'i-lucide-shuffle', title: 'Shuffle Questions', description: 'Draw random questions filtered by category or difficulty level.' },
  { icon: 'i-lucide-layers', title: '5 Categories', description: 'Icebreakers, deep dives, dreams, relationships and more.' },
  { icon: 'i-lucide-heart', title: 'Live Sessions', description: 'Create a session, draw questions one by one, and mark them answered.' }
]
</script>

<template>
  <UContainer>
    <!-- Hero -->
    <div class="py-20 text-center space-y-6">
      <UBadge color="primary" variant="subtle" size="lg">Beta</UBadge>
      <h1 class="text-5xl font-bold">
        Spark real conversations
      </h1>
      <p class="text-xl text-muted max-w-xl mx-auto">
        DateDeck shuffles thoughtful questions to help you and your date connect beyond small talk.
      </p>
      <div class="flex gap-4 justify-center">
        <UButton to="/sessions/new" size="xl" icon="i-lucide-heart">
          Start a Session
        </UButton>
        <UButton to="/questions" size="xl" variant="outline" icon="i-lucide-eye">
          Browse Questions
        </UButton>
      </div>
    </div>

    <!-- Features -->
    <div class="grid grid-cols-1 md:grid-cols-3 gap-6 pb-20">
      <UCard v-for="f in features" :key="f.title">
        <div class="flex flex-col items-center text-center gap-3 p-4">
          <UIcon :name="f.icon" class="text-primary size-10" />
          <h3 class="font-semibold text-lg">{{ f.title }}</h3>
          <p class="text-muted text-sm">{{ f.description }}</p>
        </div>
      </UCard>
    </div>
  </UContainer>
</template>
