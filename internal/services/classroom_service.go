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

type ClassroomService struct{ db *pgxpool.Pool }

func NewClassroomService(db *pgxpool.Pool) *ClassroomService { return &ClassroomService{db: db} }

func (s *ClassroomService) List(ctx context.Context, q string, year, term, page, size int) ([]models.Classroom, int, error) {
	if page <= 0 { page = 1 }
	if size <= 0 || size > 200 { size = 20 }
	offset := (page - 1) * size

	args := []any{}
	where := []string{}
	if q != "" {
		args = append(args, "%"+strings.ToLower(q)+"%")
		where = append(where, "(LOWER(name) LIKE $1 OR LOWER(code) LIKE $1)")
	}
	if year > 0 {
		args = append(args, year)
		where = append(where, fmt.Sprintf("year=$%d", len(args)))
	}
	if term > 0 {
		args = append(args, term)
		where = append(where, fmt.Sprintf("term=$%d", len(args)))
	}
	t := ""
	if len(where) > 0 { t = "WHERE " + strings.Join(where, " AND ") }

	sqlList := fmt.Sprintf(`SELECT id, code, name, grade, year, term, homeroom_teacher, capacity, created_at, updated_at
		FROM classrooms %s ORDER BY year DESC, term DESC, name ASC LIMIT %d OFFSET %d`, t, size, offset)
	rows, err := s.db.Query(ctx, sqlList, args...)
	if err != nil { return nil, 0, err }
	defer rows.Close()

	var items []models.Classroom
	for rows.Next() {
		var it models.Classroom
		if err := rows.Scan(&it.ID, &it.Code, &it.Name, &it.Grade, &it.Year, &it.Term, &it.HomeroomTeacher, &it.Capacity, &it.CreatedAt, &it.UpdatedAt); err != nil {
			return nil, 0, err
		}
		items = append(items, it)
	}
	var total int
	if err := s.db.QueryRow(ctx, "SELECT COUNT(*) FROM classrooms "+t, args...).Scan(&total); err != nil { return nil, 0, err }
	return items, total, nil
}

func (s *ClassroomService) Get(ctx context.Context, id int64) (*models.Classroom, error) {
	var it models.Classroom
	err := s.db.QueryRow(ctx, `SELECT id, code, name, grade, year, term, homeroom_teacher, capacity, created_at, updated_at FROM classrooms WHERE id=$1`, id).
		Scan(&it.ID, &it.Code, &it.Name, &it.Grade, &it.Year, &it.Term, &it.HomeroomTeacher, &it.Capacity, &it.CreatedAt, &it.UpdatedAt)
	if err == pgx.ErrNoRows { return nil, nil }
	return &it, err
}

func (s *ClassroomService) Create(ctx context.Context, it *models.Classroom) error {
	if it.Code == "" || it.Name == "" || it.Year == 0 || it.Term == 0 {
		return errors.New("missing required fields")
	}
	return s.db.QueryRow(ctx, `INSERT INTO classrooms (code,name,grade,year,term,homeroom_teacher,capacity)
		VALUES ($1,$2,$3,$4,$5,$6,COALESCE($7,50)) RETURNING id`,
		it.Code, it.Name, it.Grade, it.Year, it.Term, it.HomeroomTeacher, it.Capacity).Scan(&it.ID)
}

func (s *ClassroomService) Update(ctx context.Context, id int64, it *models.Classroom) error {
	cmd, err := s.db.Exec(ctx, `UPDATE classrooms SET name=$1, grade=$2, year=$3, term=$4, homeroom_teacher=$5, capacity=$6, updated_at=now() WHERE id=$7`,
		it.Name, it.Grade, it.Year, it.Term, it.HomeroomTeacher, it.Capacity, id)
	if err != nil { return err }
	if cmd.RowsAffected() == 0 { return pgx.ErrNoRows }
	return nil
}

func (s *ClassroomService) Delete(ctx context.Context, id int64) error {
	cmd, err := s.db.Exec(ctx, `DELETE FROM classrooms WHERE id=$1`, id)
	if err != nil { return err }
	if cmd.RowsAffected() == 0 { return pgx.ErrNoRows }
	return nil
}
