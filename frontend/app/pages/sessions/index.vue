<script setup lang="ts">
import type { CreateSessionPayload, Difficulty } from '~/types'

definePageMeta({ title: 'Sessions' })

const { sessions, loading, fetchAll, create, remove } = useSessions()
const { categories, fetchAll: fetchCategories } = useCategories()

const router = useRouter()
const showCreate = ref(false)

const form = reactive<CreateSessionPayload>({
  name: '',
  draw_count: 5,
  category_id: undefined,
  difficulty: undefined
})

const difficultyOptions: Array<{ label: string, value: Difficulty | undefined }> = [
  { label: 'Any difficulty', value: undefined },
  { label: 'Easy', value: 'easy' },
  { label: 'Medium', value: 'medium' },
  { label: 'Hard', value: 'hard' }
]

const categoryOptions = computed(() => [
  { label: 'Any category', value: undefined },
  ...categories.value.map(c => ({ label: c.name, value: c.id }))
])

onMounted(async () => {
  await fetchAll()
  await fetchCategories()
})

async function handleCreate() {
  const session = await create({ ...form })
  showCreate.value = false
  router.push(`/sessions/${session.uuid}`)
}

function formatDate(dateStr: string) {
  return new Date(dateStr).toLocaleDateString(undefined, { month: 'short', day: 'numeric', year: 'numeric' })
}
</script>

<template>
  <UContainer class="py-10 space-y-6">
    <div class="flex items-center justify-between">
      <h1 class="text-3xl font-bold">Sessions</h1>
      <UButton icon="i-lucide-plus" @click="showCreate = true">New Session</UButton>
    </div>

    <UModal v-model:open="showCreate" title="Create Session">
      <template #body>
        <UForm class="space-y-4" @submit.prevent="handleCreate">
          <UFormField label="Session name (optional)">
            <UInput v-model="form.name" placeholder="e.g. First date" class="w-full" />
          </UFormField>
          <UFormField label="Number of questions to draw">
            <UInput v-model.number="form.draw_count" type="number" min="1" max="20" class="w-full" />
          </UFormField>
          <UFormField label="Category">
            <USelect v-model="form.category_id" :items="categoryOptions" value-key="value" label-key="label" class="w-full" />
          </UFormField>
          <UFormField label="Difficulty">
            <USelect v-model="form.difficulty" :items="difficultyOptions" value-key="value" label-key="label" class="w-full" />
          </UFormField>
          <div class="flex justify-end gap-2">
            <UButton variant="ghost" @click="showCreate = false">Cancel</UButton>
            <UButton type="submit" icon="i-lucide-heart">Start</UButton>
          </div>
        </UForm>
      </template>
    </UModal>

    <div v-if="loading" class="flex justify-center py-20">
      <UIcon name="i-lucide-loader-circle" class="animate-spin size-8 text-primary" />
    </div>

    <div v-else class="space-y-3">
      <UCard
        v-for="session in sessions"
        :key="session.id"
        class="hover:ring-primary/50 cursor-pointer"
        @click="$router.push(`/sessions/${session.uuid}`)"
      >
        <div class="flex items-center justify-between p-1">
          <div class="space-y-1">
            <p class="font-semibold">{{ session.name || `Session #${session.id}` }}</p>
            <p class="text-sm text-muted">{{ formatDate(session.created_at) }}</p>
          </div>
          <div class="flex items-center gap-2">
            <UButton
              icon="i-lucide-arrow-right"
              variant="ghost"
              size="sm"
              @click.stop="$router.push(`/sessions/${session.uuid}`)"
            />
            <UButton
              icon="i-lucide-trash-2"
              color="error"
              variant="ghost"
              size="sm"
              @click.stop="remove(session.uuid)"
            />
          </div>
        </div>
      </UCard>

      <div v-if="sessions.length === 0" class="text-center py-20 space-y-4">
        <UIcon name="i-lucide-heart" class="size-12 text-muted mx-auto" />
        <p class="text-muted">No sessions yet. Start one!</p>
        <UButton icon="i-lucide-plus" @click="showCreate = true">New Session</UButton>
      </div>
    </div>
  </UContainer>
</template>
