package middleware

import "github.com/gin-gonic/gin"

// NoCache 中间件 告诉Nginx禁止缓存
// Why? https://serverfault.com/questions/801628/for-server-sent-events-sse-what-nginx-proxy-configuration-is-appropriate
func NoCache(c *gin.Context) {
	c.Header("X-Accel-Buffering", "no")
}
