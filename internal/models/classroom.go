package models

import "time"

type Classroom struct {
	ID              int64     `db:"id" json:"id"`
	Code            string    `db:"code" json:"code"`
	Name            string    `db:"name" json:"name"`
	Grade           *string   `db:"grade" json:"grade,omitempty"`
	Year            int       `db:"year" json:"year"`
	Term            int       `db:"term" json:"term"`
	HomeroomTeacher *string   `db:"homeroom_teacher" json:"homeroom_teacher,omitempty"`
	Capacity        int       `db:"capacity" json:"capacity"`
	CreatedAt       time.Time `db:"created_at" json:"created_at"`
	UpdatedAt       time.Time `db:"updated_at" json:"updated_at"`
}
