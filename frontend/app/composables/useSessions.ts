import type { CreateSessionPayload, DrawMorePayload, Session, SessionQuestion } from '~/types'

export function useSessions() {
  const api = useApi()

  const sessions = ref<Session[]>([])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      sessions.value = await api.get<Session[]>('/sessions')
    }
    catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'Failed to load sessions'
    }
    finally {
      loading.value = false
    }
  }

  async function fetchOne(uuid: string) {
    return api.get<Session>(`/sessions/uuid/${uuid}`)
  }

  async function create(payload: CreateSessionPayload) {
    const session = await api.post<Session>('/sessions', payload)
    sessions.value.unshift(session)
    return session
  }

  async function drawMore(uuid: string, payload: DrawMorePayload) {
    return api.post<Session>(`/sessions/${uuid}/draw`, payload)
  }

  async function answerQuestion(uuid: string, sqId: number) {
    return api.patch<SessionQuestion>(`/sessions/${uuid}/questions/${sqId}/answer`)
  }

  async function remove(uuid: string) {
    await api.del(`/sessions/${uuid}`)
    sessions.value = sessions.value.filter(s => s.uuid !== uuid)
  }

  return { sessions, loading, error, fetchAll, fetchOne, create, drawMore, answerQuestion, remove }
}
