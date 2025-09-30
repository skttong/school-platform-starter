package services

import (
	"context"
	"errors"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/models"
)

type AttendanceService struct{ db *pgxpool.Pool }

func NewAttendanceService(db *pgxpool.Pool) *AttendanceService { return &AttendanceService{db: db} }

func (s *AttendanceService) Record(ctx context.Context, att *models.Attendance) error {
	if att.StudentID == 0 || att.Session == "" { return errors.New("missing required fields") }
	if att.Date.IsZero() { att.Date = time.Now() }
	return s.db.QueryRow(ctx, `
		INSERT INTO attendances (student_id, classroom_id, date, session, status, note)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (student_id,classroom_id,date,session)
		DO UPDATE SET status=$5, note=$6, recorded_at=now()
		RETURNING id
	`, att.StudentID, att.ClassroomID, att.Date, att.Session, att.Status, att.Note).Scan(&att.ID)
}

func (s *AttendanceService) ListByDate(ctx context.Context, date time.Time, classroomID *int64, session string) ([]models.Attendance, error) {
	rows, err := s.db.Query(ctx, `SELECT id, student_id, classroom_id, date, session, status, note, recorded_at FROM attendances WHERE date=$1 AND session=$2 AND ($3::bigint IS NULL OR classroom_id=$3) ORDER BY student_id`, date, session, classroomID)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Attendance
	for rows.Next() { var a models.Attendance; if err := rows.Scan(&a.ID,&a.StudentID,&a.ClassroomID,&a.Date,&a.Session,&a.Status,&a.Note,&a.RecordedAt); err != nil { return nil, err }; out = append(out, a) }
	return out, nil
}


func (s *AttendanceService) SummaryDaily(ctx context.Context, date time.Time, session string) (map[string]any, error) {
	rows, err := s.db.Query(ctx, `SELECT status, COUNT(*) FROM attendances WHERE date=$1 AND session=$2 GROUP BY status`, date, session)
	if err != nil { return nil, err }
	defer rows.Close()
	total := int64(0)
	counts := map[string]int64{"PRESENT":0,"ABSENT":0,"LATE":0,"LEAVE":0}
	for rows.Next(){ var st string; var c int64; if err := rows.Scan(&st,&c); err!=nil { return nil, err }; counts[st]+=c; total+=c }
	pct := func(n int64) float64 { if total==0 { return 0 }; return float64(n)*100/float64(total) }
	return map[string]any{
		"date": date.Format("2006-01-02"),
		"session": session,
		"total": total,
		"counts": counts,
		"percent": map[string]float64{
			"PRESENT": pct(counts["PRESENT"]),
			"ABSENT": pct(counts["ABSENT"]),
			"LATE": pct(counts["LATE"]),
			"LEAVE": pct(counts["LEAVE"]),
		},
	}, nil
}

func (s *AttendanceService) SummaryClassroom(ctx context.Context, date time.Time, classroomID int64, session string) (map[string]any, error) {
	rows, err := s.db.Query(ctx, `SELECT status, COUNT(*) FROM attendances WHERE date=$1 AND session=$2 AND classroom_id=$3 GROUP BY status`, date, session, classroomID)
	if err != nil { return nil, err }
	defer rows.Close()
	total := int64(0)
	counts := map[string]int64{"PRESENT":0,"ABSENT":0,"LATE":0,"LEAVE":0}
	for rows.Next(){ var st string; var c int64; if err := rows.Scan(&st,&c); err!=nil { return nil, err }; counts[st]+=c; total+=c }
	pct := func(n int64) float64 { if total==0 { return 0 }; return float64(n)*100/float64(total) }
	return map[string]any{
		"date": date.Format("2006-01-02"),
		"session": session,
		"classroom_id": classroomID,
		"total": total,
		"counts": counts,
		"percent": map[string]float64{
			"PRESENT": pct(counts["PRESENT"]),
			"ABSENT": pct(counts["ABSENT"]),
			"LATE": pct(counts["LATE"]),
			"LEAVE": pct(counts["LEAVE"]),
		},
	}, nil
}
