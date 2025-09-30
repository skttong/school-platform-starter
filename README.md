# School Platform Starter (Go + Gin + Postgres + JWT + RBAC)

โมโนรีโประบบโรงเรียนดิจิทัล: Students / Classrooms / Enrollments / Attendance พร้อม Auth + RBAC, Swagger, Postman, Tests, CI

## Quickstart
```bash
cp .env.example .env
docker-compose up -d
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/school?sslmode=disable"

make migrate-up
go run ./cmd/hash Admin@123   # เอา hash ไปแทนใน migrations/0002_seed_rbac.sql
make seed
psql "$DATABASE_URL" -f migrations/0003_students_classes.sql
psql "$DATABASE_URL" -f migrations/0004_seed_permissions.sql
psql "$DATABASE_URL" -f migrations/0005_attendance.sql
psql "$DATABASE_URL" -f migrations/0006_seed_attendance.sql

go run ./cmd/server
# http://localhost:8080/docs
```

## Attendance (เช็คชื่อ)
- `POST /api/attendances` (record SCHOOL/CLASS)
- `GET /api/attendances?date=YYYY-MM-DD&session=CLASS&classroom_id=1`

## Make targets
`fmt`, `vet`, `lint`, `sec`, `test`, `ci`, `hook-install`


## Attendance Reports
- `GET /api/reports/attendance/daily?date=YYYY-MM-DD&session=SCHOOL` — สรุปทั้งโรงเรียน (counts + percent)
- `GET /api/reports/attendance/classroom?date=YYYY-MM-DD&session=CLASS&classroom_id=1` — สรุปเฉพาะห้อง


## Attendance Reports (เพิ่ม)
- Weekly: `GET /api/reports/attendance/weekly?start=YYYY-MM-DD&session=SCHOOL|CLASS[&classroom_id=]`
- Monthly: `GET /api/reports/attendance/monthly?year=YYYY&month=MM&session=SCHOOL|CLASS[&classroom_id=]`
- Top Absence: `GET /api/reports/attendance/top-absence?start=YYYY-MM-DD&end=YYYY-MM-DD&limit=10[&classroom_id=]`
- Export: CSV `/api/reports/attendance/export/csv?...` และ Excel `/api/reports/attendance/export/xlsx?...`
