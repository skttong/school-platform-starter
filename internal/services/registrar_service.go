package services

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/models"
)

type RegistrarService struct{ db *pgxpool.Pool }

func NewRegistrarService(db *pgxpool.Pool) *RegistrarService { return &RegistrarService{db: db} }

func (s *RegistrarService) Enroll(ctx context.Context, studentID, classroomID int64, year, term int) (*models.Enrollment, error) {
	if studentID == 0 || classroomID == 0 || year == 0 || term == 0 {
		return nil, errors.New("missing required fields")
	}
	var id int64
	err := s.db.QueryRow(ctx, `INSERT INTO enrollments (student_id, classroom_id, year, term) VALUES ($1,$2,$3,$4) RETURNING id`,
		studentID, classroomID, year, term).Scan(&id)
	if err != nil { return nil, err }
	return s.Get(ctx, id)
}

func (s *RegistrarService) UpdateStatus(ctx context.Context, id int64, status string) error {
	cmd, err := s.db.Exec(ctx, `UPDATE enrollments SET status=$1 WHERE id=$2`, status, id)
	if err != nil { return err }
	if cmd.RowsAffected() == 0 { return pgx.ErrNoRows }
	return nil
}

func (s *RegistrarService) ListByClassroom(ctx context.Context, classroomID int64, year, term int) ([]models.Enrollment, error) {
	rows, err := s.db.Query(ctx, `SELECT id, student_id, classroom_id, year, term, status, enrolled_at
		FROM enrollments WHERE classroom_id=$1 AND year=$2 AND term=$3 ORDER BY id DESC`, classroomID, year, term)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Enrollment
	for rows.Next() {
		var e models.Enrollment
		if err := rows.Scan(&e.ID, &e.StudentID, &e.ClassroomID, &e.Year, &e.Term, &e.Status, &e.EnrolledAt); err != nil { return nil, err }
		out = append(out, e)
	}
	return out, nil
}

func (s *RegistrarService) ListByStudent(ctx context.Context, studentID int64) ([]models.Enrollment, error) {
	rows, err := s.db.Query(ctx, `SELECT id, student_id, classroom_id, year, term, status, enrolled_at
		FROM enrollments WHERE student_id=$1 ORDER BY year DESC, term DESC`, studentID)
	if err != nil { return nil, err }
	defer rows.Close()
	var out []models.Enrollment
	for rows.Next() {
		var e models.Enrollment
		if err := rows.Scan(&e.ID, &e.StudentID, &e.ClassroomID, &e.Year, &e.Term, &e.Status, &e.EnrolledAt); err != nil { return nil, err }
		out = append(out, e)
	}
	return out, nil
}

func (s *RegistrarService) Get(ctx context.Context, id int64) (*models.Enrollment, error) {
	var e models.Enrollment
	err := s.db.QueryRow(ctx, `SELECT id, student_id, classroom_id, year, term, status, enrolled_at FROM enrollments WHERE id=$1`, id).
		Scan(&e.ID, &e.StudentID, &e.ClassroomID, &e.Year, &e.Term, &e.Status, &e.EnrolledAt)
	if err == pgx.ErrNoRows { return nil, nil }
	return &e, err
}
