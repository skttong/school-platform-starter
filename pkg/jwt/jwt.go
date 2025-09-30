package jwt

import (
	"time"
	jwtv5 "github.com/golang-jwt/jwt/v5"
)

type Manager struct {
	secret []byte
	iss    string
	expH   time.Duration
}

func New(secret, iss string, expireH int) *Manager {
	return &Manager{secret: []byte(secret), iss: iss, expH: time.Duration(expireH) * time.Hour}
}

func (m *Manager) Sign(sub string, extra map[string]any) (string, error) {
	claims := jwtv5.MapClaims{"sub": sub, "iss": m.iss, "iat": time.Now().Unix(), "exp": time.Now().Add(m.expH).Unix()}
	for k, v := range extra { claims[k] = v }
	t := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, claims)
	return t.SignedString(m.secret)
}

func (m *Manager) Verify(token string) (jwtv5.MapClaims, error) {
	t, err := jwtv5.Parse(token, func(t *jwtv5.Token) (any, error) { return m.secret, nil })
	if err != nil || !t.Valid { return nil, err }
	if claims, ok := t.Claims.(jwtv5.MapClaims); ok { return claims, nil }
	return nil, jwtv5.ErrTokenMalformed
}
