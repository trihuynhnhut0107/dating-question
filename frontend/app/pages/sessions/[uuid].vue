<script setup lang="ts">
import type { BadgeColor, Difficulty, DrawMorePayload, Session } from '~/types'

definePageMeta({ title: 'Session' })

const route = useRoute()
const sessionUuid = route.params.uuid as string

const { fetchOne, drawMore, answerQuestion } = useSessions()
const { categories, fetchAll: fetchCategories } = useCategories()

const session = ref<Session | null>(null)
const loading = ref(true)
const showDraw = ref(false)
const currentIndex = ref(0)

const drawForm = reactive<DrawMorePayload>({ count: 3, category_id: undefined, difficulty: undefined })

const categoryOptions = computed(() => [
  { label: 'Any category', value: undefined },
  ...categories.value.map(c => ({ label: c.name, value: c.id }))
])

const difficultyOptions: Array<{ label: string, value: Difficulty | undefined }> = [
  { label: 'Any difficulty', value: undefined },
  { label: 'Easy', value: 'easy' },
  { label: 'Medium', value: 'medium' },
  { label: 'Hard', value: 'hard' }
]

// ── derived state ──────────────────────────────────────────────────────────
const questions = computed(() => session.value?.session_questions ?? [])
const currentQuestion = computed(() => questions.value[currentIndex.value] ?? null)
const answeredCount = computed(() => questions.value.filter(sq => sq.is_answered).length)
const totalCount = computed(() => questions.value.length)
const progress = computed(() => totalCount.value > 0 ? Math.round((answeredCount.value / totalCount.value) * 100) : 0)

const difficultyColor: Partial<Record<Difficulty, BadgeColor>> = { easy: 'success', medium: 'warning', hard: 'error' }

// ── drag / swipe state ─────────────────────────────────────────────────────
const SWIPE_THRESHOLD = 90
const dragX = ref(0)
const isDragging = ref(false)
const startX = ref(0)
const isAnimating = ref(false)

const cardStyle = computed(() => ({
  transform: `translateX(${dragX.value}px) rotate(${dragX.value * 0.07}deg)`,
  transition: isDragging.value ? 'none' : 'transform 0.35s cubic-bezier(0.25, 0.46, 0.45, 0.94)'
}))

const rightOpacity = computed(() => Math.min(Math.max(dragX.value / SWIPE_THRESHOLD, 0), 1))
const leftOpacity = computed(() => Math.min(Math.max(-dragX.value / SWIPE_THRESHOLD, 0), 1))

// ── pointer handlers ───────────────────────────────────────────────────────
function onPointerDown(e: PointerEvent) {
  if (isAnimating.value || !currentQuestion.value) return
  isDragging.value = true
  startX.value = e.clientX
  dragX.value = 0
  ;(e.currentTarget as HTMLElement).setPointerCapture(e.pointerId)
}

function onPointerMove(e: PointerEvent) {
  if (!isDragging.value) return
  dragX.value = e.clientX - startX.value
}

async function onPointerUp() {
  if (!isDragging.value) return
  isDragging.value = false

  if (dragX.value > SWIPE_THRESHOLD) {
    await swipeRight()
  }
  else if (dragX.value < -SWIPE_THRESHOLD) {
    await swipeLeft()
  }
  else {
    dragX.value = 0
  }
}

async function swipeRight() {
  isAnimating.value = true
  dragX.value = 500
  await sleep(320)
  const sq = currentQuestion.value
  if (sq && !sq.is_answered) await handleAnswer(sq.id)
  advance()
  dragX.value = 0
  isAnimating.value = false
}

async function swipeLeft() {
  isAnimating.value = true
  dragX.value = -500
  await sleep(320)
  advance()
  dragX.value = 0
  isAnimating.value = false
}

// ── navigation ─────────────────────────────────────────────────────────────
function advance() {
  if (currentIndex.value < totalCount.value - 1) currentIndex.value++
  else currentIndex.value = totalCount.value
}

function goTo(i: number) { currentIndex.value = i }
function goNext() { if (currentIndex.value < totalCount.value - 1) currentIndex.value++ }
function goPrev() { if (currentIndex.value > 0) currentIndex.value-- }

// ── lifecycle ──────────────────────────────────────────────────────────────
onMounted(async () => {
  await fetchCategories()
  session.value = await fetchOne(sessionUuid)
  loading.value = false
})

async function handleAnswer(sqId: number) {
  const updated = await answerQuestion(session.value!.uuid, sqId)
  if (session.value?.session_questions) {
    const idx = session.value.session_questions.findIndex(sq => sq.id === sqId)
    if (idx !== -1) session.value.session_questions[idx] = updated
  }
}

async function handleDrawMore() {
  session.value = await drawMore(session.value!.uuid, { ...drawForm })
  showDraw.value = false
  currentIndex.value = 0
}

function sleep(ms: number) { return new Promise<void>(r => setTimeout(r, ms)) }
</script>

<template>
  <UContainer class="py-10 max-w-lg">
    <!-- Header -->
    <div class="flex items-center gap-3 mb-6">
      <UButton icon="i-lucide-arrow-left" variant="ghost" to="/sessions" />
      <h1 class="text-2xl font-bold truncate">
        {{ session?.name || `Session #${session?.id}` }}
      </h1>
    </div>

    <div v-if="loading" class="flex justify-center py-32">
      <UIcon name="i-lucide-loader-circle" class="animate-spin size-10 text-primary" />
    </div>

    <template v-else-if="session">
      <!-- Progress bar -->
      <div class="space-y-1 mb-8">
        <div class="flex justify-between text-sm text-muted">
          <span>{{ answeredCount }} / {{ totalCount }} answered</span>
          <span v-if="currentIndex < totalCount">Card {{ currentIndex + 1 }} of {{ totalCount }}</span>
        </div>
        <UProgress :value="progress" color="primary" />
      </div>

      <!-- Card stack -->
      <div class="relative h-80 mb-6 touch-none">
        <!-- Background stacked cards -->
        <div
          v-if="questions[currentIndex + 2]"
          class="absolute inset-x-4 inset-y-0 rounded-2xl bg-neutral-200 dark:bg-neutral-700"
          style="transform: translateY(10px) scale(0.90); z-index: 1;"
        />
        <div
          v-if="questions[currentIndex + 1]"
          class="absolute inset-x-2 inset-y-0 rounded-2xl bg-neutral-100 dark:bg-neutral-800 shadow"
          style="transform: translateY(5px) scale(0.95); z-index: 2;"
        />

        <!-- Active swipeable card -->
        <div
          v-if="currentQuestion"
          class="absolute inset-0 rounded-2xl shadow-2xl flex flex-col p-7 z-10 cursor-grab active:cursor-grabbing select-none overflow-hidden"
          :class="currentQuestion.is_answered ? 'bg-neutral-50 dark:bg-neutral-800' : 'bg-white dark:bg-neutral-900'"
          :style="cardStyle"
          @pointerdown="onPointerDown"
          @pointermove="onPointerMove"
          @pointerup="onPointerUp"
          @pointercancel="onPointerUp"
        >
          <!-- Green overlay (swipe right → answer) -->
          <div
            class="absolute inset-0 rounded-2xl flex items-center justify-start pl-8 pointer-events-none"
            :style="{ opacity: rightOpacity, background: 'rgba(34,197,94,0.15)' }"
          >
            <div class="flex items-center gap-2 text-green-500 font-bold text-xl border-2 border-green-500 rounded-xl px-3 py-1 -rotate-12">
              <UIcon name="i-lucide-check" class="size-5" /> DONE
            </div>
          </div>

          <!-- Red overlay (swipe left → skip) -->
          <div
            class="absolute inset-0 rounded-2xl flex items-center justify-end pr-8 pointer-events-none"
            :style="{ opacity: leftOpacity, background: 'rgba(239,68,68,0.12)' }"
          >
            <div class="flex items-center gap-2 text-red-400 font-bold text-xl border-2 border-red-400 rounded-xl px-3 py-1 rotate-12">
              SKIP <UIcon name="i-lucide-arrow-right" class="size-5" />
            </div>
          </div>

          <!-- Badges -->
          <div class="flex gap-2 flex-wrap mb-auto">
            <UBadge
              v-if="currentQuestion.question?.category"
              variant="subtle"
              size="sm"
              :style="{ background: currentQuestion.question.category.color + '22', color: currentQuestion.question.category.color }"
            >
              {{ currentQuestion.question.category.name }}
            </UBadge>
            <UBadge
              v-if="currentQuestion.question"
              :color="difficultyColor[currentQuestion.question.difficulty] ?? 'neutral'"
              variant="subtle"
              size="sm"
            >
              {{ currentQuestion.question.difficulty }}
            </UBadge>
            <UBadge v-if="currentQuestion.is_answered" color="success" variant="subtle" size="sm" icon="i-lucide-check">
              Answered
            </UBadge>
          </div>

          <!-- Question text -->
          <p class="text-xl font-semibold leading-relaxed text-center py-6">
            {{ currentQuestion.question?.text }}
          </p>

          <!-- Hint -->
          <div class="flex justify-between text-xs text-muted mt-auto">
            <span class="flex items-center gap-1">
              <UIcon name="i-lucide-arrow-left" class="size-3" /> skip
            </span>
            <span class="flex items-center gap-1">
              answer <UIcon name="i-lucide-arrow-right" class="size-3" />
            </span>
          </div>
        </div>

        <!-- All done state -->
        <div
          v-else
          class="absolute inset-0 rounded-2xl bg-white dark:bg-neutral-900 shadow-lg flex flex-col items-center justify-center gap-3 z-10"
        >
          <UIcon name="i-lucide-party-popper" class="size-14 text-primary" />
          <p class="text-xl font-semibold">All caught up!</p>
          <p class="text-sm text-muted">Draw more questions to keep going</p>
        </div>
      </div>

      <!-- Action buttons -->
      <div class="flex items-center justify-between mb-5">
        <UButton icon="i-lucide-chevron-left" variant="outline" size="sm" :disabled="currentIndex === 0" @click="goPrev">
          Prev
        </UButton>

        <UButton
          v-if="currentQuestion && !currentQuestion.is_answered"
          icon="i-lucide-check"
          color="success"
          variant="soft"
          @click="swipeRight"
        >
          Mark Done
        </UButton>
        <div v-else class="w-28" />

        <UButton icon="i-lucide-chevron-right" variant="outline" size="sm" :disabled="currentIndex >= totalCount - 1" @click="goNext">
          Next
        </UButton>
      </div>

      <!-- Dot navigator -->
      <div class="flex justify-center gap-1.5 mb-8 flex-wrap px-4">
        <button
          v-for="(sq, i) in questions"
          :key="sq.id"
          class="h-2 rounded-full transition-all duration-200"
          :class="[
            i === currentIndex
              ? 'w-5 bg-primary'
              : sq.is_answered
                ? 'w-2 bg-green-400'
                : 'w-2 bg-neutral-300 dark:bg-neutral-600'
          ]"
          @click="goTo(i)"
        />
      </div>

      <!-- Draw more -->
      <div class="flex justify-center">
        <UButton icon="i-lucide-shuffle" variant="outline" @click="showDraw = true">
          Draw More Questions
        </UButton>
      </div>

      <!-- Draw modal -->
      <UModal v-model:open="showDraw" title="Draw More Questions">
        <template #body>
          <UForm class="space-y-4" @submit.prevent="handleDrawMore">
            <UFormField label="Number of questions">
              <UInput v-model.number="drawForm.count" type="number" min="1" max="10" class="w-full" />
            </UFormField>
            <UFormField label="Category">
              <USelect v-model="drawForm.category_id" :items="categoryOptions" value-key="value" label-key="label" class="w-full" />
            </UFormField>
            <UFormField label="Difficulty">
              <USelect v-model="drawForm.difficulty" :items="difficultyOptions" value-key="value" label-key="label" class="w-full" />
            </UFormField>
            <div class="flex justify-end gap-2">
              <UButton variant="ghost" @click="showDraw = false">Cancel</UButton>
              <UButton type="submit" icon="i-lucide-shuffle">Draw</UButton>
            </div>
          </UForm>
        </template>
      </UModal>
    </template>
  </UContainer>
</template>
