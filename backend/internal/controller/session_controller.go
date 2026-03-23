package controller

import (
	"errors"

	"dating-question-backend/internal/models"
	"dating-question-backend/internal/response"
	"dating-question-backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type SessionController struct {
	svc *service.SessionService
}

func NewSessionController(svc *service.SessionService) *SessionController {
	return &SessionController{svc: svc}
}

func (ctrl *SessionController) GetAll(c *gin.Context) {
	sessions, err := ctrl.svc.GetAll()
	if err != nil {
		response.InternalError(c, "Failed to fetch sessions", err.Error())
		return
	}
	response.OK(c, "Sessions retrieved successfully", sessions)
}

func (ctrl *SessionController) GetByUUID(c *gin.Context) {
	uuid := c.Param("uuid")
	if uuid == "" {
		response.BadRequest(c, "UUID is required", "")
		return
	}
	session, err := ctrl.svc.GetByUUID(uuid)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Session not found")
			return
		}
		response.InternalError(c, "Failed to fetch session", err.Error())
		return
	}
	response.OK(c, "Session retrieved successfully", session)
}

func (ctrl *SessionController) Create(c *gin.Context) {
	var input service.CreateSessionInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	session, err := ctrl.svc.Create(input)
	if err != nil {
		response.InternalError(c, "Failed to create session", err.Error())
		return
	}
	response.Created(c, "Session created successfully", session)
}

func (ctrl *SessionController) DrawMore(c *gin.Context) {
	uuid := c.Param("uuid")

	var body struct {
		Count      int               `json:"count"`
		CategoryID *uint             `json:"category_id"`
		Difficulty *models.Difficulty `json:"difficulty"`
	}
	if err := c.ShouldBindJSON(&body); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	count := body.Count
	if count <= 0 {
		count = 5
	}

	session, err := ctrl.svc.DrawMore(uuid, count, body.CategoryID, body.Difficulty)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Session not found")
			return
		}
		response.BadRequest(c, "Could not draw more questions", err.Error())
		return
	}
	response.OK(c, "Questions drawn successfully", session)
}

func (ctrl *SessionController) AnswerQuestion(c *gin.Context) {
	uuid := c.Param("uuid")
	sqID, err := parseID(c, "sq_id")
	if err != nil {
		return
	}
	sq, err := ctrl.svc.AnswerQuestion(uuid, sqID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Session question not found")
			return
		}
		response.InternalError(c, "Failed to answer question", err.Error())
		return
	}
	response.OK(c, "Question marked as answered", sq)
}

func (ctrl *SessionController) Delete(c *gin.Context) {
	uuid := c.Param("uuid")
	if err := ctrl.svc.Delete(uuid); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Session not found")
			return
		}
		response.InternalError(c, "Failed to delete session", err.Error())
		return
	}
	response.NoContent(c)
}
