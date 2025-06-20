package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jeanyichenli/FileUploadSystem/api/handlers"
)

func SetRouter(router *gin.Engine) {

	router.POST("/upload", handlers.Upload)

}
