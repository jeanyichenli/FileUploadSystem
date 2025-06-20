package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Upload(c *gin.Context) {
	c.JSON(http.StatusOK, "upload successfully!")
}
