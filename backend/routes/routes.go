package routes

import (
	"time"

	"dating-question/controllers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	println("Registering routes with robust CORS configuration...")
	// CORS configuration
	r.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	api := r.Group("/api")
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