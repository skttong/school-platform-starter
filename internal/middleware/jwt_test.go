package middleware_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"school/internal/middleware"
	"school/pkg/jwt"
)

func TestAuthRequired(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	jm := jwt.New("secret", "issuer", 1)
	token, _ := jm.Sign("user@example.com", map[string]any{"perms": map[string]bool{"PING": true}})

	r.GET("/ping", middleware.AuthRequired(jm), func(c *gin.Context) { c.String(200, "pong") })

	req, _ := http.NewRequest("GET", "/ping", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	if w.Code != 200 || w.Body.String() != "pong" { t.Fatalf("unexpected: %d %s", w.Code, w.Body.String()) }
}
