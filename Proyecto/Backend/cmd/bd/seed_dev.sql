-- ============================================
-- SEED DEV - Datos de Desarrollo/Testing
-- ============================================
-- Este script contiene datos de prueba para desarrollo local.
-- NO debe ejecutarse en producción.

-- ============================================
-- 1. USUARIOS DE PRUEBA
-- ============================================
INSERT INTO "user" ("name", "email", "password", "role_id", "created_at")
VALUES
  -- Profesores
  ('Dr. Juan Pérez', 'juan.perez@university.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'professor'), NOW()),
  ('Dra. María González', 'maria.gonzalez@university.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'professor'), NOW()),

  -- Estudiantes
  ('Carlos Rodríguez', 'carlos.rodriguez@student.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'student'), NOW()),
  ('Ana Martínez', 'ana.martinez@student.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'student'), NOW()),
  ('Luis Fernández', 'luis.fernandez@student.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'student'), NOW()),
  ('Sofia López', 'sofia.lopez@student.edu', 'password123',
   (SELECT id FROM "role" WHERE name = 'student'), NOW())
ON CONFLICT ("email") DO NOTHING;

-- ============================================
-- 2. PROYECTO DE PRUEBA
-- ============================================
INSERT INTO "project" ("name", "objective", "created_by", "created_at", "updated_at")
VALUES (
  'Proyecto Demo - Sistema de Gestión',
  'Desarrollar un sistema web para la gestión de proyectos académicos con seguimiento de hitos y entregas.',
  (SELECT id FROM "user" WHERE email = 'carlos.rodriguez@student.edu'),
  NOW(),
  NOW()
)
ON CONFLICT DO NOTHING;

-- ============================================
-- 3. MIEMBROS DEL PROYECTO
-- ============================================
INSERT INTO "project_member" ("project_id", "user_id", "role", "joined_at")
SELECT
  p.id,
  u.id,
  CASE
    WHEN u.email = 'carlos.rodriguez@student.edu' THEN 'leader'
    ELSE 'member'
  END,
  NOW()
FROM "project" p
CROSS JOIN "user" u
WHERE p.name = 'Proyecto Demo - Sistema de Gestión'
  AND u.email IN (
    'carlos.rodriguez@student.edu',
    'ana.martinez@student.edu',
    'luis.fernandez@student.edu',
    'sofia.lopez@student.edu'
  )
ON CONFLICT DO NOTHING;

-- ============================================
-- 4. HITOS (MILESTONES) DE EJEMPLO
-- ============================================
INSERT INTO "milestone" ("project_id", "title", "description", "class_week", "created_at")
SELECT
  p.id,
  'Entrega 1 - Análisis y Diseño',
  'Documento de análisis de requisitos y diseño arquitectónico del sistema.',
  3,
  NOW()
FROM "project" p
WHERE p.name = 'Proyecto Demo - Sistema de Gestión'
ON CONFLICT DO NOTHING;

INSERT INTO "milestone" ("project_id", "title", "description", "class_week", "created_at")
SELECT
  p.id,
  'Entrega 2 - Implementación Backend',
  'Desarrollo de la API REST y modelos de datos.',
  7,
  NOW()
FROM "project" p
WHERE p.name = 'Proyecto Demo - Sistema de Gestión'
ON CONFLICT DO NOTHING;

INSERT INTO "milestone" ("project_id", "title", "description", "class_week", "created_at")
SELECT
  p.id,
  'Entrega 3 - Frontend y Testing',
  'Desarrollo de la interfaz de usuario y pruebas integradas.',
  12,
  NOW()
FROM "project" p
WHERE p.name = 'Proyecto Demo - Sistema de Gestión'
ON CONFLICT DO NOTHING;

-- ============================================
-- 5. COMENTARIOS DE EJEMPLO
-- ============================================
INSERT INTO "comment" ("milestone_id", "user_id", "content", "created_at")
SELECT
  m.id,
  u.id,
  'Gran avance en esta entrega. El diseño arquitectónico está bien estructurado.',
  NOW() - INTERVAL '2 days'
FROM "milestone" m
CROSS JOIN "user" u
WHERE m.title = 'Entrega 1 - Análisis y Diseño'
  AND u.email = 'ana.martinez@student.edu'
LIMIT 1
ON CONFLICT DO NOTHING;

-- ============================================
-- Fin de datos de desarrollo
-- ============================================
