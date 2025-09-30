package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"school/internal/models"
	"school/internal/services"
)

type StudentHandler struct{ svc *services.StudentService }

func NewStudentHandler(s *services.StudentService) *StudentHandler { return &StudentHandler{svc: s} }

func (h *StudentHandler) List(c *gin.Context) {
	q := c.Query("q")
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	size, _ := strconv.Atoi(c.DefaultQuery("page_size", "20"))
	items, total, err := h.svc.List(c, q, services.Page{Page: page, PageSize: size})
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"items": items, "total": total, "page": page, "page_size": size})
}

func (h *StudentHandler) Get(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	st, err := h.svc.Get(c, id)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	if st == nil { c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return }
	c.JSON(http.StatusOK, st)
}

func (h *StudentHandler) Create(c *gin.Context) {
	var req models.Student
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
	if err := h.svc.Create(c, &req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, req)
}

func (h *StudentHandler) Update(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	var req models.Student
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"}); return }
	if err := h.svc.Update(c, id, &req); err != nil {
		if err.Error() == "no rows in result set" { c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return }
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	c.Status(http.StatusNoContent)
}

func (h *StudentHandler) Delete(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	if err := h.svc.Delete(c, id); err != nil {
		if err.Error() == "no rows in result set" { c.JSON(http.StatusNotFound, gin.H{"error": "not found"}); return }
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return
	}
	c.Status(http.StatusNoContent)
}
