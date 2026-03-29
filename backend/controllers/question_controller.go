package controllers

import (
	"dating-question/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func ListQuestion(c *gin.Context){
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	filter:= services.QuestionFilter{
		Cursor: c.Query("cursor"),
		Limit: limit,
		CategoryID: c.Query("category_id"),
		Search:c.Query("search"),
		Sort:c.Query("sort"),
	}
	questions, nextCursor, err := services.GetQuestionsByCursor(filter)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":        questions,
		"next_cursor": nextCursor,
		"limit":       limit,
	})
}