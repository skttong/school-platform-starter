package services

import (
	"context"
	"crypto/sha256"
	"encoding/hex"

	"github.com/jackc/pgx/v5/pgxpool"
)

type AuthService struct{ db *pgxpool.Pool }

func NewAuthService(db *pgxpool.Pool) *AuthService { return &AuthService{db: db} }

func (s *AuthService) Hash(pw string) string {
	h := sha256.Sum256([]byte(pw))
	return hex.EncodeToString(h[:])
}

func (s *AuthService) ValidateUser(ctx context.Context, email, password string) (userID int64, perms map[string]bool, ok bool) {
	var id int64
	var hash string
	if err := s.db.QueryRow(ctx, `SELECT id, password FROM users WHERE email=$1 AND is_active=true`, email).Scan(&id, &hash); err != nil { return 0, nil, false }
	if hash != s.Hash(password) { return 0, nil, false }
	rows, _ := s.db.Query(ctx, `
		SELECT p.code FROM permissions p
		JOIN role_permissions rp ON rp.permission_id=p.id
		JOIN user_roles ur ON ur.role_id=rp.role_id
		WHERE ur.user_id=$1`, id)
	perms = map[string]bool{}
	for rows.Next() { var code string; _ = rows.Scan(&code); perms[code] = true }
	return id, perms, true
}
