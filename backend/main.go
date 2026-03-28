package main

import (
	"dating-question/config"
	"dating-question/routes"

	"github.com/gin-gonic/gin"
)
func main() {
	config.ConnectDatabase()
	r:= gin.Default()
	routes.RegisterRoutes(r)
	r.Run(":8080")
}