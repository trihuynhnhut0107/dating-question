package controllers

import (
	"dating-question/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var input struct {
		Name           *string `json:"name"`
		QuestionNumber *int    `json:"question_number"`
	}

	ip := c.ClientIP()

	if err := c.ShouldBindJSON(&input); err != nil {
		// skip error for optional name
	}

	session, err := services.CreateSession(input.Name, &ip, input.QuestionNumber)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, session)
}
