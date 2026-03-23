package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type ErrResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   string `json:"error,omitempty"`
}

func OK(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, Response{Success: true, Message: message, Data: data})
}

func Created(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusCreated, Response{Success: true, Message: message, Data: data})
}

func BadRequest(c *gin.Context, message string, err string) {
	c.JSON(http.StatusBadRequest, ErrResponse{Success: false, Message: message, Error: err})
}

func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, ErrResponse{Success: false, Message: message})
}

func Conflict(c *gin.Context, message string, err string) {
	c.JSON(http.StatusConflict, ErrResponse{Success: false, Message: message, Error: err})
}

func InternalError(c *gin.Context, message string, err string) {
	c.JSON(http.StatusInternalServerError, ErrResponse{Success: false, Message: message, Error: err})
}

func NoContent(c *gin.Context) {
	c.Status(http.StatusNoContent)
}
