import type { QuestionResponseDto } from "./question";

export interface CreateSessionDto {
  name?: string;
  question_number?: number;
}

export interface SessionResponseDto {
  id: string;
  name: string;
  questions: QuestionResponseDto[];
  created_at: string;
  updated_at: string;
}
