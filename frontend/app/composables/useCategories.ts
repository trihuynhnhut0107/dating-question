import type { Category, CreateCategoryPayload } from '~/types'

export function useCategories() {
  const api = useApi()

  const categories = useState<Category[]>('categories', () => [])
  const loading = ref(false)
  const error = ref<string | null>(null)

  async function fetchAll() {
    loading.value = true
    error.value = null
    try {
      categories.value = await api.get<Category[]>('/categories')
    }
    catch (e: unknown) {
      error.value = e instanceof Error ? e.message : 'Failed to load categories'
    }
    finally {
      loading.value = false
    }
  }

  async function create(payload: CreateCategoryPayload) {
    const category = await api.post<Category>('/categories', payload)
    categories.value.push(category)
    return category
  }

  async function remove(id: number) {
    await api.del(`/categories/${id}`)
    categories.value = categories.value.filter(c => c.id !== id)
  }

  return { categories, loading, error, fetchAll, create, remove }
}
