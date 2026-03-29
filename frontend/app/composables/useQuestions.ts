import type { CategoryDto } from "~/types/category";
import type { PaginatedResponse } from "~/types/common";
import type { QuestionDto } from "~/types/question";

const CATEGORIES_ENDPOINT = "/categories";
const QUESTIONS_ENDPOINT = "/questions";

export const useQuestions = () => {
  const fetchCategories = () => useApi<CategoryDto[]>(CATEGORIES_ENDPOINT);
  const fetchQuestionsByCursor = (params: {
    cursor?: string;
    limit?: number;
    category_id?: string;
    search?: string;
    sort?: string;
  }) => {
    const query = new URLSearchParams();
    if (params.cursor) query.append("cursor", params.cursor);
    if (params.limit) query.append("limit", params.limit.toString());
    if (params.category_id) query.append("category_id", params.category_id);
    if (params.search) query.append("search", params.search);
    if (params.sort) query.append("sort", params.sort);
    return useApi<PaginatedResponse<QuestionDto>>(QUESTIONS_ENDPOINT);
  };
  return {
    fetchCategories,
    fetchQuestionsByCursor,
  };
};
