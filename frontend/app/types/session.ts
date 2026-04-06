import type { QuestionDto } from "./question";

export interface CreateSessionDto {
  name: string;
}

export interface SessionDto {
  id: string;
  name: string;
  questions: QuestionDto[];
  createdAt: Date;
  updatedAt: Date;
}
