package middlewares

import (
	"fmt"
	"webapi-with-go/services"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		const Bearer_schema = "Bearer "
		header := c.GetHeader("Authorization")
		fmt.Println(header)
		if header == "" {
			c.AbortWithStatus(401)
		}

		token := header[len(Bearer_schema):]

		if !services.NewJWTService().ValidadeToken(token) {
			c.AbortWithStatus(401)
		}
	}
}
