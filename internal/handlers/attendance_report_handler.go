package handlers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"school/internal/services"
)

type AttendanceReportHandler struct{ svc *services.AttendanceService }

func NewAttendanceReportHandler(s *services.AttendanceService) *AttendanceReportHandler { return &AttendanceReportHandler{svc: s} }

func (h *AttendanceReportHandler) Daily(c *gin.Context) {
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	d, _ := time.Parse("2006-01-02", dateStr)
	session := c.DefaultQuery("session", "SCHOOL")
	res, err := h.svc.SummaryDaily(c, d, session)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, res)
}

func (h *AttendanceReportHandler) Classroom(c *gin.Context) {
	dateStr := c.DefaultQuery("date", time.Now().Format("2006-01-02"))
	d, _ := time.Parse("2006-01-02", dateStr)
	session := c.DefaultQuery("session", "CLASS")
	cidStr := c.Query("classroom_id")
	if cidStr == "" { c.JSON(http.StatusBadRequest, gin.H{"error":"classroom_id required"}); return }
	cid, err := strconv.ParseInt(cidStr, 10, 64); if err != nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid classroom_id"}); return }
	res, err := h.svc.SummaryClassroom(c, d, cid, session)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }
	c.JSON(http.StatusOK, res)
}
