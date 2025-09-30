package jwt_test

import (
	"testing"
	"time"
	"school/pkg/jwt"
)

func TestJWTSignVerify(t *testing.T) {
	m := jwt.New("secret","issuer",1)
	tok, err := m.Sign("user@example.com", map[string]any{"uid": int64(1)})
	if err != nil { t.Fatalf("sign: %v", err) }
	claims, err := m.Verify(tok)
	if err != nil { t.Fatalf("verify: %v", err) }
	if claims["sub"].(string) != "user@example.com" { t.Fatalf("sub mismatch") }
	if float64(time.Now().Unix()) > claims["exp"].(float64) { t.Fatalf("expired") }
}
