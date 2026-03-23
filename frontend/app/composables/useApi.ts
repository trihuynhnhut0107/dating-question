import type { ApiResponse, QueryParams } from '~/types'

export function useApi() {
  const config = useRuntimeConfig()
  const baseURL = config.public.apiBase

  async function get<T>(path: string, query?: QueryParams): Promise<T> {
    const params = query
      ? Object.fromEntries(Object.entries(query).filter(([, v]) => v !== undefined).map(([k, v]) => [k, String(v)]))
      : undefined
    const res = await $fetch<ApiResponse<T>>(`${baseURL}${path}`, { query: params })
    return res.data
  }

  async function post<T>(path: string, body?: unknown): Promise<T> {
    const res = await $fetch<ApiResponse<T>>(`${baseURL}${path}`, { method: 'POST', body: body as Record<string, unknown> })
    return res.data
  }

  async function put<T>(path: string, body?: unknown): Promise<T> {
    const res = await $fetch<ApiResponse<T>>(`${baseURL}${path}`, { method: 'PUT', body: body as Record<string, unknown> })
    return res.data
  }

  async function patch<T>(path: string, body?: unknown): Promise<T> {
    const res = await $fetch<ApiResponse<T>>(`${baseURL}${path}`, { method: 'PATCH', body: body as Record<string, unknown> })
    return res.data
  }

  async function del(path: string): Promise<void> {
    await $fetch(`${baseURL}${path}`, { method: 'DELETE' })
  }

  return { get, post, put, patch, del }
}
