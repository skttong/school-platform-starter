package services

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"school/internal/models"
)

type StudentService struct{ db *pgxpool.Pool }

func NewStudentService(db *pgxpool.Pool) *StudentService { return &StudentService{db: db} }

type Page struct{ Page, PageSize int }

func (s *StudentService) List(ctx context.Context, q string, p Page) ([]models.Student, int, error) {
	if p.Page <= 0 { p.Page = 1 }
	if p.PageSize <= 0 || p.PageSize > 200 { p.PageSize = 20 }
	offset := (p.Page - 1) * p.PageSize
	args := []any{}
	where := []string{}
	if q != "" { args = append(args, "%"+strings.ToLower(q)+"%"); where = append(where, "(LOWER(first_name)||' '||LOWER(last_name) LIKE $1 OR LOWER(student_code) LIKE $1)") }
	w := ""; if len(where) > 0 { w = "WHERE "+strings.Join(where, " AND ") }
	sqlList := fmt.Sprintf(`SELECT id, student_code, prefix, first_name, last_name, gender, birth_date, citizen_id,
		phone, email, address, guardian_name, guardian_phone, status, created_at, updated_at
		FROM students %s ORDER BY id DESC LIMIT %d OFFSET %d`, w, p.PageSize, offset)
	rows, err := s.db.Query(ctx, sqlList, args...); if err != nil { return nil, 0, err }
	defer rows.Close()
	var items []models.Student
	for rows.Next() { var st models.Student; if err := rows.Scan(&st.ID,&st.StudentCode,&st.Prefix,&st.FirstName,&st.LastName,&st.Gender,&st.BirthDate,&st.CitizenID,&st.Phone,&st.Email,&st.Address,&st.GuardianName,&st.GuardianPhone,&st.Status,&st.CreatedAt,&st.UpdatedAt); err != nil { return nil,0,err }; items = append(items, st) }
	var total int; if err := s.db.QueryRow(ctx, "SELECT COUNT(*) FROM students "+w, args...).Scan(&total); err != nil { return nil,0,err }
	return items, total, nil
}

func (s *StudentService) Get(ctx context.Context, id int64) (*models.Student, error) {
	var st models.Student
	err := s.db.QueryRow(ctx, `SELECT id, student_code, prefix, first_name, last_name, gender, birth_date, citizen_id,
		phone, email, address, guardian_name, guardian_phone, status, created_at, updated_at FROM students WHERE id=$1`, id).
		Scan(&st.ID,&st.StudentCode,&st.Prefix,&st.FirstName,&st.LastName,&st.Gender,&st.BirthDate,&st.CitizenID,&st.Phone,&st.Email,&st.Address,&st.GuardianName,&st.GuardianPhone,&st.Status,&st.CreatedAt,&st.UpdatedAt)
	if err == pgx.ErrNoRows { return nil, nil }
	return &st, err
}

func (s *StudentService) Create(ctx context.Context, st *models.Student) error {
	if st.StudentCode == "" || st.FirstName == "" || st.LastName == "" { return errors.New("missing required fields") }
	return s.db.QueryRow(ctx, `INSERT INTO students (student_code,prefix,first_name,last_name,gender,birth_date,citizen_id,phone,email,address,guardian_name,guardian_phone,status)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,COALESCE($13,'ACTIVE')) RETURNING id`,
		st.StudentCode, st.Prefix, st.FirstName, st.LastName, st.Gender, st.BirthDate, st.CitizenID, st.Phone, st.Email, st.Address, st.GuardianName, st.GuardianPhone, st.Status).Scan(&st.ID)
}

func (s *StudentService) Update(ctx context.Context, id int64, st *models.Student) error {
	cmd, err := s.db.Exec(ctx, `UPDATE students SET prefix=$1, first_name=$2, last_name=$3, gender=$4, birth_date=$5, citizen_id=$6, phone=$7, email=$8, address=$9, guardian_name=$10, guardian_phone=$11, status=$12, updated_at=now() WHERE id=$13`,
		st.Prefix, st.FirstName, st.LastName, st.Gender, st.BirthDate, st.CitizenID, st.Phone, st.Email, st.Address, st.GuardianName, st.GuardianPhone, st.Status, id)
	if err != nil { return err }
	if cmd.RowsAffected() == 0 { return pgx.ErrNoRows }
	return nil
}

func (s *StudentService) Delete(ctx context.Context, id int64) error {
	cmd, err := s.db.Exec(ctx, `DELETE FROM students WHERE id=$1`, id); if err != nil { return err }
	if cmd.RowsAffected() == 0 { return pgx.ErrNoRows }
	return nil
}
