package jwt_test

import (
	"testing"
	"time"

	"school/pkg/jwt"
)

func TestJWTSignVerify(t *testing.T) {
	m := jwt.New("secret", "issuer", 1) // 1 hour
	tok, err := m.Sign("user@example.com", map[string]any{"uid": int64(123)})
	if err != nil { t.Fatalf("sign error: %v", err) }
	claims, err := m.Verify(tok)
	if err != nil { t.Fatalf("verify error: %v", err) }
	if claims["sub"].(string) != "user@example.com" { t.Fatalf("unexpected sub: %v", claims["sub"]) }
	if claims["iss"].(string) != "issuer" { t.Fatalf("unexpected iss: %v", claims["iss"]) }
	if _, ok := claims["uid"]; !ok { t.Fatalf("missing uid claim") }
	// expire check (not exact, but exp must be in the future)
	if float64(time.Now().Unix()) > claims["exp"].(float64) { t.Fatalf("token already expired") }
}
