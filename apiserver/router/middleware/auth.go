/**
 * 功能描述: API身份认证
 * @Date: 2019-04-17
 * @author: lixiaoming
 */
package middleware

import (
	"apiserver/controllers"
	"apiserver/pkg/errno"
	"apiserver/pkg/token"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if _, err := token.ParseRequest(c); err != nil {
			controllers.SendResponse(c, errno.ErrTokenInvaild, nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
