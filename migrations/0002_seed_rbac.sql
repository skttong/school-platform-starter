INSERT INTO roles (code, name) VALUES
  ('ADMIN','System Admin'),
  ('TEACHER','Teacher'),
  ('REGISTRAR','Registrar')
ON CONFLICT DO NOTHING;

INSERT INTO permissions (code, name) VALUES
  ('USER_READ','Read Users'),
  ('USER_WRITE','Write Users'),
  ('STUDENT_READ','Read Students'),
  ('STUDENT_WRITE','Write Students')
ON CONFLICT DO NOTHING;

-- map ADMIN to all current permissions
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r, permissions p
WHERE r.code='ADMIN'
ON CONFLICT DO NOTHING;

-- demo user: admin@example.com / Admin@123 (sha256 DEMO ให้เปลี่ยน)
INSERT INTO users (email, password, full_name)
VALUES ('admin@example.com', 'REPLACE_WITH_SHA256_FROM_cmd_hash', 'System Admin')
ON CONFLICT DO NOTHING;

INSERT INTO user_roles(user_id, role_id)
SELECT u.id, r.id FROM users u, roles r
WHERE u.email='admin@example.com' AND r.code='ADMIN'
ON CONFLICT DO NOTHING;
