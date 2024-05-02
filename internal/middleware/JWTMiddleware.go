package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/tadilbek11kz/ePharma-backend/internal/app/config"
	"github.com/tadilbek11kz/ePharma-backend/internal/util/token"
)

type JWTMiddleware struct {
}

func NewJWTMiddleware() *JWTMiddleware {
	return &JWTMiddleware{}
}

func (m *JWTMiddleware) New() gin.HandlerFunc {
	return func(c *gin.Context) {
		var access_token string
		cookie, err := c.Cookie("access_token")

		authorizationHeader := c.Request.Header.Get("Authorization")
		fields := strings.Fields(authorizationHeader)

		if len(fields) != 0 && fields[0] == "Bearer" {
			access_token = fields[1]
		} else if err == nil {
			access_token = cookie
		}

		if access_token == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": "You are not logged in"})
			return
		}

		config := config.NewConfig()
		sub, err := token.ValidateToken(access_token, config.AccessTokenPublicKey)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"status": "fail", "message": err.Error()})
			return
		}

		c.Set("user", sub)
		c.Next()
	}
}
