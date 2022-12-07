package middleware

import (
	"github.com/ibnumei/go-ms-playground/internal/app/domain"
	"context"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
)

func WithAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		if !strings.HasPrefix(authHeader, "Bearer ") {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		auths := strings.Split(authHeader, " ")
		if len(auths) != 2 {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
			return
		}

		//decrypt jwt
		user := domain.User{}
		data, err := user.DecryptJWT(auths[1])
		fmt.Println(data)
		if err != nil {
			c.JSON(401, map[string]string{
				"message": "unauthorized",
			})
			c.Abort()
		}
		ctxUserID := context.WithValue(c.Request.Context(), "user_id", int(data["user_id"].(float64)))
		c.Request = c.Request.WithContext(ctxUserID)
		
		c.Next()
	}
}