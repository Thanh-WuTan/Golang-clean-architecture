package middleware

import (
	"net/http"
	common "onlyfounds/common"
	"onlyfounds/internal/tokenutil"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtAuthMiddleware(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization")
		t := strings.Split(authHeader, " ")
		if len(t) == 2 && t[0] == "Bearer" && strings.TrimSpace(t[1]) != "" {
			authToken := t[1]
			authorized, err := tokenutil.IsAuthorized(authToken, secret)
			if authorized {
				username, err := tokenutil.ExtractUserNameFromToken(authToken, secret)
				if err != nil {
					c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: err.Error()})
					c.Abort()
					return
				}
				c.Set("current_user", username)
				c.Next()
				return
			}
			c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: err.Error()})
			c.Abort()
			return
		}
		c.JSON(http.StatusUnauthorized, common.ErrorResponse{Message: "Not authorized"})
		c.Abort()
	}
}
