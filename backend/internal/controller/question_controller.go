package controller

import (
	"errors"
	"strconv"

	"dating-question-backend/internal/models"
	"dating-question-backend/internal/response"
	"dating-question-backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type QuestionController struct {
	svc *service.QuestionService
}

func NewQuestionController(svc *service.QuestionService) *QuestionController {
	return &QuestionController{svc: svc}
}

func (ctrl *QuestionController) GetAll(c *gin.Context) {
	filter := service.ListQuestionsFilter{}

	if catID := c.Query("category_id"); catID != "" {
		if id, err := strconv.ParseUint(catID, 10, 64); err == nil {
			v := uint(id)
			filter.CategoryID = &v
		}
	}
	if diff := c.Query("difficulty"); diff != "" {
		d := models.Difficulty(diff)
		filter.Difficulty = &d
	}
	if active := c.Query("is_active"); active != "" {
		v := active == "true"
		filter.IsActive = &v
	}

	questions, err := ctrl.svc.GetAll(filter)
	if err != nil {
		response.InternalError(c, "Failed to fetch questions", err.Error())
		return
	}
	response.OK(c, "Questions retrieved successfully", questions)
}

func (ctrl *QuestionController) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	question, err := ctrl.svc.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Question not found")
			return
		}
		response.InternalError(c, "Failed to fetch question", err.Error())
		return
	}
	response.OK(c, "Question retrieved successfully", question)
}

func (ctrl *QuestionController) Create(c *gin.Context) {
	var input service.CreateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	question, err := ctrl.svc.Create(input)
	if err != nil {
		response.InternalError(c, "Failed to create question", err.Error())
		return
	}
	response.Created(c, "Question created successfully", question)
}

func (ctrl *QuestionController) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	var input service.UpdateQuestionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	question, err := ctrl.svc.Update(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Question not found")
			return
		}
		response.InternalError(c, "Failed to update question", err.Error())
		return
	}
	response.OK(c, "Question updated successfully", question)
}

func (ctrl *QuestionController) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	if err := ctrl.svc.Delete(id); err != nil {
		response.InternalError(c, "Failed to delete question", err.Error())
		return
	}
	response.NoContent(c)
}

func (ctrl *QuestionController) GetRandom(c *gin.Context) {
	count := 5
	if n := c.Query("count"); n != "" {
		if parsed, err := strconv.Atoi(n); err == nil && parsed > 0 {
			count = parsed
		}
	}

	var categoryID *uint
	if catID := c.Query("category_id"); catID != "" {
		if id, err := strconv.ParseUint(catID, 10, 64); err == nil {
			v := uint(id)
			categoryID = &v
		}
	}

	var difficulty *models.Difficulty
	if diff := c.Query("difficulty"); diff != "" {
		d := models.Difficulty(diff)
		difficulty = &d
	}

	questions, err := ctrl.svc.GetRandom(count, categoryID, difficulty)
	if err != nil {
		response.InternalError(c, "Failed to fetch random questions", err.Error())
		return
	}
	response.OK(c, "Random questions retrieved successfully", questions)
}
