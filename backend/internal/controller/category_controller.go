package controller

import (
	"errors"
	"strconv"

	"dating-question-backend/internal/response"
	"dating-question-backend/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type CategoryController struct {
	svc *service.CategoryService
}

func NewCategoryController(svc *service.CategoryService) *CategoryController {
	return &CategoryController{svc: svc}
}

func (ctrl *CategoryController) GetAll(c *gin.Context) {
	categories, err := ctrl.svc.GetAll()
	if err != nil {
		response.InternalError(c, "Failed to fetch categories", err.Error())
		return
	}
	response.OK(c, "Categories retrieved successfully", categories)
}

func (ctrl *CategoryController) GetByID(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	category, err := ctrl.svc.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Category not found")
			return
		}
		response.InternalError(c, "Failed to fetch category", err.Error())
		return
	}
	response.OK(c, "Category retrieved successfully", category)
}

func (ctrl *CategoryController) Create(c *gin.Context) {
	var input service.CreateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	category, err := ctrl.svc.Create(input)
	if err != nil {
		response.Conflict(c, "Failed to create category", err.Error())
		return
	}
	response.Created(c, "Category created successfully", category)
}

func (ctrl *CategoryController) Update(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	var input service.UpdateCategoryInput
	if err := c.ShouldBindJSON(&input); err != nil {
		response.BadRequest(c, "Invalid request body", err.Error())
		return
	}
	category, err := ctrl.svc.Update(id, input)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "Category not found")
			return
		}
		response.InternalError(c, "Failed to update category", err.Error())
		return
	}
	response.OK(c, "Category updated successfully", category)
}

func (ctrl *CategoryController) Delete(c *gin.Context) {
	id, err := parseID(c, "id")
	if err != nil {
		return
	}
	if err := ctrl.svc.Delete(id); err != nil {
		response.InternalError(c, "Failed to delete category", err.Error())
		return
	}
	response.NoContent(c)
}

// parseID is a shared helper within the controller package.
func parseID(c *gin.Context, param string) (uint, error) {
	raw, err := strconv.ParseUint(c.Param(param), 10, 64)
	if err != nil {
		response.BadRequest(c, "Invalid ID", "ID must be a positive integer")
		return 0, err
	}
	return uint(raw), nil
}
