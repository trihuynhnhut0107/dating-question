import type { CategoryResponseDto } from "~/types/category";
import type { PaginatedResponse } from "~/types/common";
import type { QuestionFilterDto, QuestionResponseDto } from "~/types/question";

const CATEGORIES_ENDPOINT = "/categories";
const QUESTIONS_ENDPOINT = "/questions";

export const useQuestions = () => {
  const fetchCategories = () =>
    useApi<CategoryResponseDto[]>(CATEGORIES_ENDPOINT);
  const fetchQuestionsByCursor = async (params: QuestionFilterDto) => {
    const data = await useApi<PaginatedResponse<QuestionResponseDto>>(
      QUESTIONS_ENDPOINT,
      {
        query: params,
      },
    );
    return (
      data.data.value ?? {
        data: [],
        limit: 10,
        next_cursor: "",
      }
    );
  };
  return {
    fetchCategories,
    fetchQuestionsByCursor,
  };
};
