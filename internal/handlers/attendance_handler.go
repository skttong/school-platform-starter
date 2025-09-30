package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"school/internal/models"
	"school/internal/services"
)

type AttendanceHandler struct{ svc *services.AttendanceService }

func NewAttendanceHandler(s *services.AttendanceService) *AttendanceHandler { return &AttendanceHandler{svc: s} }

func (h *AttendanceHandler) Record(c *gin.Context) {
	var req models.Attendance
	if err := c.ShouldBindJSON(&req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid payload"}); return }
	if err := h.svc.Record(c, &req); err != nil { c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusCreated, req)
}

func (h *AttendanceHandler) ListByDate(c *gin.Context) {
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	d, _ := time.Parse("2006-01-02", dateStr)
	session := c.DefaultQuery("session","SCHOOL")
	var classID *int64
	if s := c.Query("classroom_id"); s != "" { if v, err := strconv.ParseInt(s,10,64); err==nil { classID = &v } }
	items, err := h.svc.ListByDate(c, d, classID, session)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, gin.H{"items": items})
}
