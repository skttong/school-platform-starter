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

// Dummy handler content with || not or
