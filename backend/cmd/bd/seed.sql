-- ============================================
-- SEED DATA - Datos Iniciales Esenciales
-- ============================================
-- Este script contiene datos base necesarios para el funcionamiento de la aplicación.
-- Se ejecuta después de init.sql y es idempotente (puede ejecutarse múltiples veces).

-- ============================================
-- 1. ROLES
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
-- NOTA: La autenticación se realiza exclusivamente mediante Google OAuth.
-- Este usuario admin es solo para propósitos de prueba o herramientas internas.
-- ProviderID: usar un ID ficticio o real de Google OAuth.
INSERT INTO "user" ("name", "email", "provider_id", "role_id", "created_at")
VALUES (
  'Administrador',
  'admin@softpharos.com',
  'google-oauth-admin-123',
  (SELECT id FROM "role" WHERE name = 'admin'),
  NOW()
)
ON CONFLICT ("email") DO NOTHING;

-- ============================================
-- Fin de datos esenciales
-- ============================================
