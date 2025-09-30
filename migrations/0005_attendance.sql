CREATE TABLE IF NOT EXISTS attendances (
  id BIGSERIAL PRIMARY KEY,
  student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
  classroom_id BIGINT REFERENCES classrooms(id) ON DELETE CASCADE,
  date DATE NOT NULL,
  session TEXT NOT NULL,
  status TEXT NOT NULL DEFAULT 'PRESENT',
  note TEXT,
  recorded_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(student_id, classroom_id, date, session)
);
CREATE INDEX IF NOT EXISTS idx_attendances_student ON attendances (student_id, date);
CREATE INDEX IF NOT EXISTS idx_attendances_classroom ON attendances (classroom_id, date);
