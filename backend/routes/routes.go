package routes

import (
	"dating-question/controllers"

	"github.com/gin-gonic/gin"
)
func RegisterRoutes(r *gin.Engine){
	api:=r.Group("/api")
	{
		v1:= api.Group("/v1")
		{
			questions:= v1.Group("/questions")
			{
				questions.GET("", controllers.ListQuestion)
			}

		}
	}
}