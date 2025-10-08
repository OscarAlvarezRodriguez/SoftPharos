package main

import (
	"context"
	"github.com/joho/godotenv"
	"log"
	"softpharos/internal/infra/databases"
	"softpharos/internal/infra/databases/models"
	"softpharos/internal/infra/repository"
)

func main() {
	// Cargar variables del archivo .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️  No se pudo cargar .env, usando variables de entorno existentes")
	}

	// Crear cliente PostgreSQL con GORM desde variables de entorno
	client, err := databases.NewClientFromEnv()
	if err != nil {
		log.Fatalf("❌ Error al conectar con PostgreSQL: %v", err)
	}
	defer client.Close()

	// Verificar conexión
	if err := client.Ping(); err != nil {
		log.Fatalf("❌ Error al hacer ping a la BD: %v", err)
	}

	log.Println("✅ Conectado a PostgreSQL con GORM")

	// Crear repositorio de roles
	roleRepo := repository.NewRoleRepository(client)

	ctx := context.Background()

	// Ejemplo 1: Obtener todos los roles
	log.Println("\n📋 Obteniendo todos los roles...")
	roles, err := roleRepo.GetAll(ctx)
	if err != nil {
		log.Printf("⚠️  Error al obtener roles: %v", err)
	} else {
		log.Printf("✅ Se encontraron %d roles", len(roles))
		for _, role := range roles {
			log.Printf("   - ID: %d, Nombre: %s", role.ID, role.Name)
		}
	}

	// Ejemplo 2: Intentar obtener un rol específico por ID (si existe)
	log.Println("\n🔍 Buscando rol con ID 1...")
	role, err := roleRepo.GetByID(ctx, 1)
	if err != nil {
		log.Printf("⚠️  No se encontró el rol con ID 1: %v", err)
	} else {
		log.Printf("✅ Rol encontrado: %s", role.Name)
		if role.Description != nil {
			log.Printf("   Descripción: %s", *role.Description)
		}
	}

	// Ejemplo 3: Buscar por nombre
	log.Println("\n🔍 Buscando rol 'admin'...")
	adminRole, err := roleRepo.GetByName(ctx, "Administrador")
	if err != nil {
		log.Printf("⚠️  No se encontró el rol 'admin': %v", err)
	} else {
		log.Printf("✅ Rol admin encontrado con ID: %d", adminRole.ID)
	}

	// Ejemplo 4: Consulta directa con GORM para ver usuarios (usando modelos de persistencia)
	log.Println("\n👥 Obteniendo todos los usuarios...")
	var users []models.UserModel
	result := client.DB.Preload("Role").Find(&users)
	if result.Error != nil {
		log.Printf("⚠️  Error al obtener usuarios: %v", result.Error)
	} else {
		log.Printf("✅ Se encontraron %d usuarios", len(users))
		for _, user := range users {
			roleName := "Sin rol"
			if user.Role != nil {
				roleName = user.Role.Name
			}
			log.Printf("   - %s (%s) - Rol: %s", user.Email, getValue(user.Name), roleName)
		}
	}

	log.Println("\n🎉 Ejemplo completado exitosamente")
}

// getValue es una función helper para obtener el valor de un puntero string
func getValue(s *string) string {
	if s == nil {
		return "Sin nombre"
	}
	return *s
}
