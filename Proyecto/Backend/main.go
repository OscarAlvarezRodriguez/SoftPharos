package main

import (
	"github.com/joho/godotenv"
	"log"
	"softpharos/internal/infra/databases"
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

	log.Println("\n🎉 Ejemplo completado exitosamente")
}
