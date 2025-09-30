package models

import "time"

type Enrollment struct {
	ID          int64     `db:"id" json:"id"`
	StudentID   int64     `db:"student_id" json:"student_id"`
	ClassroomID int64     `db:"classroom_id" json:"classroom_id"`
	Year        int       `db:"year" json:"year"`
	Term        int       `db:"term" json:"term"`
	Status      string    `db:"status" json:"status"`
	EnrolledAt  time.Time `db:"enrolled_at" json:"enrolled_at"`
}
