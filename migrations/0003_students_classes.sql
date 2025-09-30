CREATE TABLE IF NOT EXISTS students (
  id BIGSERIAL PRIMARY KEY,
  student_code TEXT UNIQUE NOT NULL,
  prefix TEXT,
  first_name TEXT NOT NULL,
  last_name TEXT NOT NULL,
  gender TEXT,
  birth_date DATE,
  citizen_id TEXT UNIQUE,
  phone TEXT,
  email TEXT,
  address TEXT,
  guardian_name TEXT,
  guardian_phone TEXT,
  status TEXT NOT NULL DEFAULT 'ACTIVE',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS classrooms (
  id BIGSERIAL PRIMARY KEY,
  code TEXT UNIQUE NOT NULL,
  name TEXT NOT NULL,
  grade TEXT,
  year INT NOT NULL,
  term INT NOT NULL,
  homeroom_teacher TEXT,
  capacity INT NOT NULL DEFAULT 50,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);
CREATE TABLE IF NOT EXISTS enrollments (
  id BIGSERIAL PRIMARY KEY,
  student_id BIGINT NOT NULL REFERENCES students(id) ON DELETE CASCADE,
  classroom_id BIGINT NOT NULL REFERENCES classrooms(id) ON DELETE CASCADE,
  year INT NOT NULL,
  term INT NOT NULL,
  status TEXT NOT NULL DEFAULT 'ENROLLED',
  enrolled_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  UNIQUE(student_id, classroom_id, year, term)
);
CREATE INDEX IF NOT EXISTS idx_students_name ON students (last_name, first_name);
CREATE INDEX IF NOT EXISTS idx_classrooms_year_term ON classrooms (year, term);
CREATE INDEX IF NOT EXISTS idx_enrollments_student ON enrollments (student_id);
