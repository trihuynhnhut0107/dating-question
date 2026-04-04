package routes

import (
	"dating-question/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(cors.Default())
	{
		v1 := api.Group("/v1")
		{
			questions := v1.Group("/questions")
			{
				questions.GET("", controllers.ListQuestion)
			}

			sessions := v1.Group("/sessions")
			{
				sessions.POST("", controllers.CreateSession)
			}
		}
	}
}