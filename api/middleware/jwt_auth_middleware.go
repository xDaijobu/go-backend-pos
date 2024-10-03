package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"go-backend-pos/domain"
	"go-backend-pos/internal/tokenutil"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				claims, err := tokenutil.ExtractClaimsFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}

				//uaString := c.Request.Header.Get("User-Agent")

				c.Set("x-user-id", claims.UserID)
				c.Set("x-session-id", claims.SessionID)
				c.Set("x-auth-token", authToken)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, domain.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
