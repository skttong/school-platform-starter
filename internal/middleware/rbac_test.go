package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"school/internal/middleware"
)

func TestRequirePermission(t *testing.T){
	gin.SetMode(gin.TestMode)
	r := gin.New()
	r.Use(func(c *gin.Context){ c.Set("perms", map[string]bool{"ALLOW": true}); c.Next() })
	r.GET("/ok", middleware.RequirePermission("ALLOW"), func(c *gin.Context){ c.String(200,"ok") })
	r.GET("/deny", middleware.RequirePermission("DENY"), func(c *gin.Context){ c.String(200,"no") })
	w1 := httptest.NewRecorder(); req1,_ := http.NewRequest("GET","/ok",nil); r.ServeHTTP(w1,req1); if w1.Code!=200 { t.Fatalf("want 200 got %d", w1.Code) }
	w2 := httptest.NewRecorder(); req2,_ := http.NewRequest("GET","/deny",nil); r.ServeHTTP(w2,req2); if w2.Code!=403 { t.Fatalf("want 403 got %d", w2.Code) }
}
