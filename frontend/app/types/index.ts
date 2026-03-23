export type Difficulty = 'easy' | 'medium' | 'hard'

export interface Category {
  id: number
  created_at: string
  updated_at: string
  name: string
  description: string
  color: string
  questions?: Question[]
}

export interface Question {
  id: number
  created_at: string
  updated_at: string
  text: string
  category_id: number
  category?: Category
  difficulty: Difficulty
  is_active: boolean
}

export interface SessionQuestion {
  id: number
  created_at: string
  updated_at: string
  session_id: number
  question_id: number
  question?: Question
  is_answered: boolean
  answered_at: string | null
}

export interface Session {
  id: number
  uuid: string
  created_at: string
  updated_at: string
  name: string
  session_questions?: SessionQuestion[]
}

export interface ApiResponse<T> {
  success: boolean
  message: string
  data: T
}

export interface ApiError {
  success: false
  message: string
  error: string
}

// Nuxt UI badge/button color literal union
export type BadgeColor = 'error' | 'primary' | 'secondary' | 'success' | 'info' | 'warning' | 'neutral'

// Shared query param type — all filter interfaces extend this so they're
// assignable to the api.get() query parameter without unsafe casts.
export type QueryParams = Record<string, string | number | boolean | undefined>

export interface QuestionsFilter extends QueryParams {
  category_id?: number
  difficulty?: Difficulty
  is_active?: boolean
}

export interface RandomQuestionsFilter extends QueryParams {
  count?: number
  category_id?: number
  difficulty?: Difficulty
}

// Request types
export interface CreateCategoryPayload {
  name: string
  description?: string
  color?: string
}

export interface CreateQuestionPayload {
  text: string
  category_id: number
  difficulty?: Difficulty
  is_active?: boolean
}

export interface CreateSessionPayload {
  name?: string
  draw_count?: number
  category_id?: number
  difficulty?: Difficulty
}

export interface DrawMorePayload {
  count?: number
  category_id?: number
  difficulty?: Difficulty
}
