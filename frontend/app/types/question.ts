export interface QuestionResponseDto {
  id: string;
  content: string;
  category_id: string;
  created_at?: string;
  updated_at?: string;
}

export interface QuestionFilterDto {
  cursor?: string;
  limit?: number;
  category_id?: string;
  session_id?: string;
  search?: string;
  sort?: string;
}

// For future use when creating questions
export interface CreateQuestionDto {
  content: string;
  category_id: string;
}
