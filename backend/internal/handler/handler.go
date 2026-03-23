package handler

import (
	"net/http"

	"dating-question-backend/internal/controller"
	"dating-question-backend/internal/db"
	"dating-question-backend/internal/response"
	"dating-question-backend/internal/service"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	// ── utility ───────────────────────────────────────────────────────────────
	router.GET("/ping", func(c *gin.Context) {
		response.OK(c, "pong", nil)
	})
	router.GET("/health", func(c *gin.Context) {
		sqlDB, err := db.DB.DB()
		if err != nil || sqlDB.Ping() != nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"success": false, "message": "database unreachable"})
			return
		}
		response.OK(c, "healthy", gin.H{"database": "ok"})
	})

	// ── dependencies ──────────────────────────────────────────────────────────
	categorySvc := service.NewCategoryService(db.DB)
	questionSvc := service.NewQuestionService(db.DB)
	sessionSvc := service.NewSessionService(db.DB)

	categoryCtrl := controller.NewCategoryController(categorySvc)
	questionCtrl := controller.NewQuestionController(questionSvc)
	sessionCtrl := controller.NewSessionController(sessionSvc)

	api := router.Group("/api/v1")
	{
		// ── categories ────────────────────────────────────────────────────────
		categories := api.Group("/categories")
		{
			categories.GET("", categoryCtrl.GetAll)
			categories.GET("/:id", categoryCtrl.GetByID)
			categories.POST("", categoryCtrl.Create)
			categories.PUT("/:id", categoryCtrl.Update)
			categories.DELETE("/:id", categoryCtrl.Delete)
		}

		// ── questions ─────────────────────────────────────────────────────────
		questions := api.Group("/questions")
		{
			questions.GET("", questionCtrl.GetAll)         // ?category_id=&difficulty=&is_active=
			questions.GET("/random", questionCtrl.GetRandom) // ?count=&category_id=&difficulty=
			questions.GET("/:id", questionCtrl.GetByID)
			questions.POST("", questionCtrl.Create)
			questions.PUT("/:id", questionCtrl.Update)
			questions.DELETE("/:id", questionCtrl.Delete)
		}

		// ── sessions ──────────────────────────────────────────────────────────
		sessions := api.Group("/sessions")
		{
			sessions.GET("", sessionCtrl.GetAll)
			sessions.GET("/:uuid", sessionCtrl.GetByUUID)
			sessions.POST("", sessionCtrl.Create)
			sessions.POST("/:uuid/draw", sessionCtrl.DrawMore)
			sessions.PATCH("/:uuid/questions/:sq_id/answer", sessionCtrl.AnswerQuestion)
			sessions.DELETE("/:uuid", sessionCtrl.Delete)
		}
	}

	return router
}
