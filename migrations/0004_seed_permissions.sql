INSERT INTO permissions (code, name) VALUES
 ('CLASSROOM_READ','Read Classrooms'),
 ('CLASSROOM_WRITE','Write Classrooms'),
 ('ENROLLMENT_READ','Read Enrollments'),
 ('ENROLLMENT_WRITE','Write Enrollments')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p WHERE r.code='ADMIN' ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p WHERE r.code='TEACHER' AND p.code IN ('STUDENT_READ','CLASSROOM_READ','ENROLLMENT_READ') ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p WHERE r.code='REGISTRAR' AND p.code IN ('CLASSROOM_READ','CLASSROOM_WRITE','ENROLLMENT_READ','ENROLLMENT_WRITE','STUDENT_READ','STUDENT_WRITE') ON CONFLICT DO NOTHING;
