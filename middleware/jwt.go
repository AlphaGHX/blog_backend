package middleware

import (
	"blog/models/response"

	"blog/utils"
	"github.com/gin-gonic/gin"
)

func JwtAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			response.FailWithMessage("未登录", c)
			c.Abort()
			return
		}
		claims, err := utils.VerifyToken(token)
		if err != nil {
			response.FailWithMessage("Token验证失败。Toklen错误或过期", c)
			c.Abort()
			return 
		}
		c.Set("claims", claims)
		c.Next()
	}
}
