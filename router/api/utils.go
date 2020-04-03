package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func handleNotFound(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{
		"status": 404,
		"error":  "404, page not ready to GO!",
	})
}
