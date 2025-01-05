package routers

import (
	"api/handlers"

	"github.com/gin-gonic/gin"
)

func Health(r *gin.RouterGroup){
	handlers := handlers.NewHealthHandler()

	r.GET("/", handlers.HealthGET)
	r.POST("/", handlers.HealthPOST)
	r.POST("/:id", handlers.HealthPOSTById)

}