<script setup lang="ts">
import type { BadgeColor, Difficulty } from '~/types'

definePageMeta({ title: 'Questions' })

const { questions, loading, fetchAll } = useQuestions()
const { categories, fetchAll: fetchCategories } = useCategories()

const filterCategoryId = ref<number | undefined>()
const filterDifficulty = ref<Difficulty | undefined>()
const filterActive = ref<boolean>(true)

const difficultyOptions: Array<{ label: string, value: Difficulty | undefined }> = [
  { label: 'All', value: undefined },
  { label: 'Easy', value: 'easy' },
  { label: 'Medium', value: 'medium' },
  { label: 'Hard', value: 'hard' }
]

const categoryOptions = computed(() => [
  { label: 'All categories', value: undefined },
  ...categories.value.map(c => ({ label: c.name, value: c.id }))
])

const difficultyColor: Partial<Record<Difficulty, BadgeColor>> = {
  easy: 'success',
  medium: 'warning',
  hard: 'error'
}

onMounted(async () => {
  await fetchCategories()
  await applyFilters()
})

async function applyFilters() {
  await fetchAll({
    category_id: filterCategoryId.value,
    difficulty: filterDifficulty.value,
    is_active: filterActive.value
  })
}

watch([filterCategoryId, filterDifficulty, filterActive], applyFilters)
</script>

<template>
  <UContainer class="py-10 space-y-6">
    <div class="flex items-center justify-between">
      <h1 class="text-3xl font-bold">Questions</h1>
      <UBadge size="lg" variant="subtle">{{ questions.length }} results</UBadge>
    </div>

    <!-- Filters -->
    <div class="flex flex-wrap gap-4 items-center">
      <USelect
        v-model="filterCategoryId"
        :items="categoryOptions"
        value-key="value"
        label-key="label"
        placeholder="Category"
        class="w-48"
      />
      <USelect
        v-model="filterDifficulty"
        :items="difficultyOptions"
        value-key="value"
        label-key="label"
        placeholder="Difficulty"
        class="w-36"
      />
      <UCheckbox v-model="filterActive" label="Active only" />
    </div>

    <div v-if="loading" class="flex justify-center py-20">
      <UIcon name="i-lucide-loader-circle" class="animate-spin size-8 text-primary" />
    </div>

    <div v-else class="space-y-3">
      <UCard v-for="q in questions" :key="q.id" class="hover:ring-primary/50">
        <div class="flex items-start justify-between gap-4 p-1">
          <div class="space-y-1">
            <p class="font-medium">{{ q.text }}</p>
            <div class="flex items-center gap-2 flex-wrap">
              <UBadge
                v-if="q.category"
                variant="subtle"
                size="sm"
                :style="{ background: q.category.color + '22', color: q.category.color }"
              >
                {{ q.category.name }}
              </UBadge>
              <UBadge :color="difficultyColor[q.difficulty] || 'neutral'" variant="subtle" size="sm">
                {{ q.difficulty }}
              </UBadge>
            </div>
          </div>
          <UIcon v-if="!q.is_active" name="i-lucide-eye-off" class="text-muted shrink-0" />
        </div>
      </UCard>

      <p v-if="questions.length === 0" class="text-center text-muted py-10">
        No questions match your filters.
      </p>
    </div>
  </UContainer>
</template>
