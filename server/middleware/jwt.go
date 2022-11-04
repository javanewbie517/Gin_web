package middleware

import (
	"bubble/server/response"
	"bubble/server/util/jwt"

	"github.com/gin-gonic/gin"
)

/**
JWTAuth
	校验token，并解析token包含的信息，然后放入context往下传递
*/

func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 我们这里jwt鉴权取头部信息 token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
		token := c.Request.Header.Get("token")
		if token == "" {
			response.FailWithDetailed(c, gin.H{"reload": true}, "未登录或非法访问")
			c.Abort()
			return
		}
		// parseToken 解析token包含的信息
		claims, err := jwt.ParseToken(token)
		if err != nil {
			response.FailWithDetailed(c, gin.H{"reload": true}, err.Error())
			c.Abort()
			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}
