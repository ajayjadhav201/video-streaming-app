package routes

import (
	"watch-service/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.GET("/videos/*filepath", controller.VideoHandler)
}
