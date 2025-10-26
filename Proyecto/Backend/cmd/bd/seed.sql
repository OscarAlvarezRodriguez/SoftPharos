-- ============================================
-- SEED DATA - Datos Iniciales Esenciales
-- ============================================
-- Este script contiene datos base necesarios para el funcionamiento de la aplicación.
-- Se ejecuta después de init.sql y es idempotente (puede ejecutarse múltiples veces).

-- ============================================
-- 1. ROLES (Datos esenciales)
-- ============================================
INSERT INTO "role" ("name", "description", "created_at")
VALUES
  ('admin', 'Administrador del sistema', NOW()),
  ('professor', 'Profesor que supervisa proyectos', NOW()),
  ('student', 'Estudiante participante en proyectos', NOW())
ON CONFLICT ("name") DO NOTHING;

-- ============================================
-- 2. USUARIO ADMINISTRADOR POR DEFECTO
-- ============================================
-- Contraseña: admin123 (CAMBIAR EN PRODUCCIÓN)
-- Nota: Hashea esta contraseña en producción usando bcrypt
INSERT INTO "user" ("name", "email", "password", "role_id", "created_at")
VALUES (
  'Administrador',
  'admin@softpharos.com',
  'admin123', -- TODO: Hashear con bcrypt en producción
  (SELECT id FROM "role" WHERE name = 'admin'),
  NOW()
)
ON CONFLICT ("email") DO NOTHING;

-- ============================================
-- Fin de datos esenciales
-- ============================================
