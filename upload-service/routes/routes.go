package routes

import (
	"upload-service/controller"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	//
	r.POST("/upload", controller.UploadHandler)
}
