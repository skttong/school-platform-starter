package models

import "time"

type Student struct {
	ID            int64      `db:"id" json:"id"`
	StudentCode   string     `db:"student_code" json:"student_code"`
	Prefix        string     `db:"prefix" json:"prefix"`
	FirstName     string     `db:"first_name" json:"first_name"`
	LastName      string     `db:"last_name" json:"last_name"`
	Gender        string     `db:"gender" json:"gender"`
	BirthDate     *time.Time `db:"birth_date" json:"birth_date,omitempty"`
	CitizenID     *string    `db:"citizen_id" json:"citizen_id,omitempty"`
	Phone         *string    `db:"phone" json:"phone,omitempty"`
	Email         *string    `db:"email" json:"email,omitempty"`
	Address       *string    `db:"address" json:"address,omitempty"`
	GuardianName  *string    `db:"guardian_name" json:"guardian_name,omitempty"`
	GuardianPhone *string    `db:"guardian_phone" json:"guardian_phone,omitempty"`
	Status        string     `db:"status" json:"status"`
	CreatedAt     time.Time  `db:"created_at" json:"created_at"`
	UpdatedAt     time.Time  `db:"updated_at" json:"updated_at"`
}
