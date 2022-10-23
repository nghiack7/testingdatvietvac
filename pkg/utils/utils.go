package utils

import (
	"github.com/gin-gonic/gin"
)

func BuildResponseWithGinError(err error, errCode int, c *gin.Context) {
	c.JSON(-1, gin.H{err.Error(): errCode})
}
