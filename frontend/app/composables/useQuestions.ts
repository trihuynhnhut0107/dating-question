import type { CreateQuestionPayload, QuestionsFilter, RandomQuestionsFilter, Question } from '~/types'

export function useQuestions() {
  const api = useApi()

  const questions = ref<Question[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll(filters?: QuestionsFilter) {
    loading.value = true
    error.value = null
    try {
      questions.value = await api.get<Question[]>('/questions', filters)
    }
    catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'Failed to load questions'
    }
    finally {
      loading.value = false
    }
  }

  async function fetchRandom(opts?: RandomQuestionsFilter) {
    loading.value = true
    try {
      return await api.get<Question[]>('/questions/random', opts)
    }
    finally {
      loading.value = false
    }
  }

  async function create(payload: CreateQuestionPayload) {
    return api.post<Question>('/questions', payload)
  }

  async function remove(id: number) {
    await api.del(`/questions/${id}`)
    questions.value = questions.value.filter(q => q.id !== id)
  }

  return { questions, loading, error, fetchAll, fetchRandom, create, remove }
}
