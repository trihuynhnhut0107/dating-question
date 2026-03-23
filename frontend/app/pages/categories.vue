<script setup lang="ts">
import type { CreateCategoryPayload } from '~/types'

definePageMeta({ title: 'Categories' })

const { categories, loading, fetchAll, create, remove } = useCategories()

const showCreate = ref(false)
const form = reactive<CreateCategoryPayload>({ name: '', description: '', color: '#6366f1' })

onMounted(fetchAll)

async function handleCreate() {
  await create({ ...form })
  showCreate.value = false
  form.name = ''
  form.description = ''
  form.color = '#6366f1'
}
</script>

<template>
  <UContainer class="py-10 space-y-6">
    <div class="flex items-center justify-between">
      <h1 class="text-3xl font-bold">Categories</h1>
      <UButton icon="i-lucide-plus" @click="showCreate = true">New Category</UButton>
    </div>

    <UModal v-model:open="showCreate" title="Create Category">
      <template #body>
        <UForm class="space-y-4" @submit.prevent="handleCreate">
          <UFormField label="Name" required>
            <UInput v-model="form.name" placeholder="e.g. Icebreaker" class="w-full" />
          </UFormField>
          <UFormField label="Description">
            <UTextarea v-model="form.description" placeholder="Short description…" class="w-full" />
          </UFormField>
          <UFormField label="Color">
            <UInput v-model="form.color" type="color" class="w-16 h-9" />
          </UFormField>
          <div class="flex justify-end gap-2">
            <UButton variant="ghost" @click="showCreate = false">Cancel</UButton>
            <UButton type="submit">Create</UButton>
          </div>
        </UForm>
      </template>
    </UModal>

    <div v-if="loading" class="flex justify-center py-20">
      <UIcon name="i-lucide-loader-circle" class="animate-spin size-8 text-primary" />
    </div>

    <div v-else class="grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-4">
      <UCard v-for="cat in categories" :key="cat.id">
        <div class="flex items-start justify-between p-1">
          <div class="flex items-center gap-3">
            <span class="size-4 rounded-full shrink-0 mt-1" :style="{ background: cat.color }" />
            <div>
              <p class="font-semibold">{{ cat.name }}</p>
              <p class="text-sm text-muted">{{ cat.description }}</p>
            </div>
          </div>
          <UButton
            icon="i-lucide-trash-2"
            color="error"
            variant="ghost"
            size="sm"
            @click="remove(cat.id)"
          />
        </div>
      </UCard>
    </div>
  </UContainer>
</template>
