# School Platform Starter (Go + Gin + Postgres + JWT + RBAC)

Starter สำหรับระบบโรงเรียนดิจิทัล พร้อมโมดูล นักเรียน / ห้องเรียน / ทะเบียน (ลงทะเบียนเรียน)

## คุณสมบัติ
- REST API + JWT Auth
- RBAC (Roles/Permissions) พร้อม seed
- โมดูล Students, Classrooms, Enrollments
- Docker Compose (Postgres + Adminer)

## เริ่มใช้งานอย่างย่อ
```bash
cp .env.example .env
docker-compose up -d
export DATABASE_URL="postgres://postgres:postgres@localhost:5432/school?sslmode=disable"
make migrate-up
psql "$DATABASE_URL" -f migrations/0002_seed_rbac.sql
psql "$DATABASE_URL" -f migrations/0003_students_classes.sql
psql "$DATABASE_URL" -f migrations/0004_seed_permissions.sql
go run ./cmd/hash Admin@123   # นำ hash ไปแทนใน 0002 seed แล้วค่อย
make seed
go run ./cmd/server
```

ดูเอกสาร endpoints เพิ่มในโค้ดที่โฟลเดอร์ `internal/handlers` และ `internal/routes`.


## API Docs (Swagger / OpenAPI)
- เปิดดูที่ `http://localhost:8080/docs`
- ไฟล์สคีมาอยู่ที่ `api/openapi.yaml`

## CI
- มี GitHub Actions (`.github/workflows/ci.yml`) สำหรับ build + test อัตโนมัติบน Go 1.22


## Postman
- อยู่ที่ `api/postman_collection.json` (import เข้า Postman แล้วแก้ตัวแปร `baseUrl` และ `token`)

## Lint & Security
- ใช้ `golangci-lint` (config `.golangci.yml`) และ `gosec` ผ่าน GitHub Actions job `lint-and-sec`


## Swagger Examples
- ตัวอย่าง request/response ถูกใส่ไว้ใน `api/openapi.yaml` ครบทั้ง list/get/post/put/delete สำหรับ Students, Classrooms และ Enrollments

## Postman Environment
- ใช้ `api/postman_local_environment.json` ร่วมกับ collection เพื่อเซ็ต `baseUrl` และ `token` เร็วขึ้น

## Pre-commit Hook
```bash
make hook-install      # ติดตั้งสคริปต์ pre-commit (ลินต์ + เทสต์ + vet + gosec)
```
> ถ้ายังไม่มี `golangci-lint` และ `gosec` ให้ติดตั้ง:
```bash
go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
go install github.com/securego/gosec/v2/cmd/gosec@latest
```

## Make targets (ใหม่)
- `make fmt` / `make vet` / `make lint` / `make sec` / `make test`
- `make ci` รวบยอด tidy + build + test + lint + sec
