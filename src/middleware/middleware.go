package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func BodySizeMaxMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		var w http.ResponseWriter = c.Writer
		//add limit =10kb
		var limit int64 = 10 * 1024
		c.Request.Body = http.MaxBytesReader(w, c.Request.Body, limit)
		c.Next()
	}

}
