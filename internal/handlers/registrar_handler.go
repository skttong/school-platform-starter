package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"school/internal/services"
)

type RegistrarHandler struct{ svc *services.RegistrarService }

func NewRegistrarHandler(s *services.RegistrarService) *RegistrarHandler { return &RegistrarHandler{svc: s} }

func (h *RegistrarHandler) Enroll(c *gin.Context) {
	var req struct{ StudentID int64 `json:"student_id"`; ClassroomID int64 `json:"classroom_id"`; Year int `json:"year"`; Term int `json:"term"` }
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid payload"}); return }
	enr, err := h.svc.Enroll(c, req.StudentID, req.ClassroomID, req.Year, req.Term)
	if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, enr)
}

func (h *RegistrarHandler) UpdateStatus(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"),10,64)
	var req struct{ Status string `json:"status"` }
	if err := c.ShouldBindJSON(&req); err != nil || req.Status=="" { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid payload"}); return }
	if err := h.svc.UpdateStatus(c, id, req.Status); err != nil { if err.Error()=="no rows in result set" { c.JSON(http.StatusNotFound, gin.H{"error":"not found"}); return }; c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.Status(http.StatusNoContent)
}

func (h *RegistrarHandler) ListByClassroom(c *gin.Context) {
	cid, _ := strconv.ParseInt(c.Param("classroom_id"),10,64)
	year, _ := strconv.Atoi(c.DefaultQuery("year","0"))
	term, _ := strconv.Atoi(c.DefaultQuery("term","0"))
	items, err := h.svc.ListByClassroom(c, cid, year, term)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"items": items})
}

func (h *RegistrarHandler) ListByStudent(c *gin.Context) {
	sid, _ := strconv.ParseInt(c.Param("student_id"),10,64)
	items, err := h.svc.ListByStudent(c, sid)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"items": items})
}
