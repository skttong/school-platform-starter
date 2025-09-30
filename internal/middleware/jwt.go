package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"school/pkg/jwt"
)

func AuthRequired(j *jwt.Manager) gin.HandlerFunc {
	return func(c *gin.Context) {
		h := c.GetHeader("Authorization")
		parts := strings.SplitN(h, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "missing bearer token"})
			return
		}
		claims, err := j.Verify(parts[1])
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Set("claims", claims)
		// optional: map perms if present
		if p, ok := claims["perms"]; ok {
			c.Set("perms", p)
		}
		c.Next()
	}
}
