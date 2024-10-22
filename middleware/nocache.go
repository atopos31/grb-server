package middleware

import "github.com/gin-gonic/gin"

func NoCache(c *gin.Context) {
	c.Header("X-Accel-Buffering","no")
}