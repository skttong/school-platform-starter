package handlers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"
	"encoding/csv"

	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
	"school/internal/services"
)

type AttendanceExportHandler struct{ svc *services.AttendanceService }

func NewAttendanceExportHandler(s *services.AttendanceService) *AttendanceExportHandler { return &AttendanceExportHandler{svc: s} }

func (h *AttendanceExportHandler) CSV(c *gin.Context) {
	startStr := c.Query("start"); endStr := c.Query("end")
	if startStr=="" || endStr=="" { c.JSON(http.StatusBadRequest, gin.H{"error":"start and end required"}); return }
	start, err1 := time.Parse("2006-01-02", startStr)
	end, err2 := time.Parse("2006-01-02", endStr)
	if err1!=nil || err2!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid date"}); return }
	session := c.DefaultQuery("session","SCHOOL")
	var classID *int64
	if cid := c.Query("classroom_id"); cid != "" { if v, err := strconv.ParseInt(cid,10,64); err==nil { classID = &v } }
	res, err := h.svc.SummaryRange(c, start, end, session, classID)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=attendance_%s_to_%s.csv", startStr, endStr))
	c.Header("Content-Type", "text/csv; charset=utf-8")
	w := csv.NewWriter(c.Writer)
	defer w.Flush()

	w.Write([]string{"date","PRESENT","ABSENT","LATE","LEAVE"})
	if daily, ok := res["daily"].(map[string]map[string]int64); ok {
		for d, ct := range daily {
			w.Write([]string{ d, fmt.Sprintf("%d", ct["PRESENT"]), fmt.Sprintf("%d", ct["ABSENT"]), fmt.Sprintf("%d", ct["LATE"]), fmt.Sprintf("%d", ct["LEAVE"]) })
		}
	} else if m2, ok := res["daily"].(map[string]interface{}); ok {
		for d, v := range m2 {
			row := v.(map[string]interface{})
			w.Write([]string{ d, fmt.Sprintf("%v", row["PRESENT"]), fmt.Sprintf("%v", row["ABSENT"]), fmt.Sprintf("%v", row["LATE"]), fmt.Sprintf("%v", row["LEAVE"]) })
		}
	}
}

func (h *AttendanceExportHandler) XLSX(c *gin.Context) {
	startStr := c.Query("start"); endStr := c.Query("end")
	if startStr=="" || endStr=="" { c.JSON(http.StatusBadRequest, gin.H{"error":"start and end required"}); return }
	start, err1 := time.Parse("2006-01-02", startStr)
	end, err2 := time.Parse("2006-01-02", endStr)
	if err1!=nil || err2!=nil { c.JSON(http.StatusBadRequest, gin.H{"error":"invalid date"}); return }
	session := c.DefaultQuery("session","SCHOOL")
	var classID *int64
	if cid := c.Query("classroom_id"); cid != "" { if v, err := strconv.ParseInt(cid,10,64); err==nil { classID = &v } }
	res, err := h.svc.SummaryRange(c, start, end, session, classID)
	if err != nil { c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()}); return }

	f := excelize.NewFile()
	sheet := "Attendance"
	f.SetSheetName("Sheet1", sheet)
	f.SetCellValue(sheet, "A1", "date"); f.SetCellValue(sheet, "B1", "PRESENT"); f.SetCellValue(sheet, "C1", "ABSENT"); f.SetCellValue(sheet, "D1", "LATE"); f.SetCellValue(sheet, "E1", "LEAVE")

	row := 2
	if daily, ok := res["daily"].(map[string]map[string]int64); ok {
		for d, ct := range daily {
			f.SetCellValue(sheet, fmt.Sprintf("A%d", row), d)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", row), ct["PRESENT"])
			f.SetCellValue(sheet, fmt.Sprintf("C%d", row), ct["ABSENT"])
			f.SetCellValue(sheet, fmt.Sprintf("D%d", row), ct["LATE"])
			f.SetCellValue(sheet, fmt.Sprintf("E%d", row), ct["LEAVE"])
			row++
		}
	} else if m2, ok := res["daily"].(map[string]interface{}); ok {
		for d, v := range m2 {
			rowm := v.(map[string]interface{})
			f.SetCellValue(sheet, fmt.Sprintf("A%d", row), d)
			f.SetCellValue(sheet, fmt.Sprintf("B%d", row), rowm["PRESENT"]) 
			f.SetCellValue(sheet, fmt.Sprintf("C%d", row), rowm["ABSENT"]) 
			f.SetCellValue(sheet, fmt.Sprintf("D%d", row), rowm["LATE"]) 
			f.SetCellValue(sheet, fmt.Sprintf("E%d", row), rowm["LEAVE"]) 
			row++
		}
	}

	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=attendance_%s_to_%s.xlsx", startStr, endStr))
	c.Header("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	_ = f.Write(c.Writer)
}
