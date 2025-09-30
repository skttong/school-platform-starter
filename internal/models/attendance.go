package models

import "time"

type Attendance struct {
	ID         int64     `db:"id" json:"id"`
	StudentID  int64     `db:"student_id" json:"student_id"`
	ClassroomID *int64   `db:"classroom_id,omitempty" json:"classroom_id,omitempty"`
	Date       time.Time `db:"date" json:"date"`
	Session    string    `db:"session" json:"session"` // SCHOOL / CLASS
	Status     string    `db:"status" json:"status"`   // PRESENT / ABSENT / LATE / LEAVE
	Note       *string   `db:"note" json:"note,omitempty"`
	RecordedAt time.Time `db:"recorded_at" json:"recorded_at"`
}
