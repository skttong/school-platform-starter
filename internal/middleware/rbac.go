package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RequirePermission(perm string) gin.HandlerFunc {
	return func(c *gin.Context) {
		perms, _ := c.Get("perms")
		ps, ok := perms.(map[string]bool)
		if !ok {
			// try map[string]interface{} from JWT claims
			if m2, ok2 := perms.(map[string]interface{}); ok2 {
				ps = map[string]bool{}
				for k, v := range m2 {
					if b, ok := v.(bool); ok && b { ps[k] = true }
				}
				ok = true
			}
		}
		if !ok || !ps[perm] {
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"error": "forbidden"})
			return
		}
		c.Next()
	}
}
