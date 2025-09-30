INSERT INTO permissions (code,name) VALUES
 ('ATTENDANCE_READ','Read Attendance'),
 ('ATTENDANCE_WRITE','Write Attendance')
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.code='ADMIN' AND p.code LIKE 'ATTENDANCE_%'
ON CONFLICT DO NOTHING;

INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.code='TEACHER' AND p.code IN ('ATTENDANCE_READ','ATTENDANCE_WRITE')
ON CONFLICT DO NOTHING;
