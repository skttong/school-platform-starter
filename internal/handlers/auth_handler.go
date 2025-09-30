package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"school/internal/services"
	"school/pkg/jwt"
)

type AuthHandler struct { auth *services.AuthService; jm *jwt.Manager }

func NewAuthHandler(auth *services.AuthService, jm *jwt.Manager) *AuthHandler { return &AuthHandler{auth: auth, jm: jm} }

func (h *AuthHandler) Login(c *gin.Context) {
	var req struct{ Email, Password string }
	if err := c.ShouldBindJSON(&req); err != nil || req.Email == "" || req.Password == "" { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid payload"}); return }
	id, perms, ok := h.auth.ValidateUser(c, req.Email, req.Password)
	if !ok { c.JSON(http.StatusUnauthorized, gin.H{"error":"invalid credentials"}); return }
	token, _ := h.jm.Sign(req.Email, map[string]any{"uid": id, "perms": perms})
	c.JSON(http.StatusOK, gin.H{"access_token": token})
}
