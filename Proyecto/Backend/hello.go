package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

func main() {
	// Cargar variables del archivo .env
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️ No se pudo cargar .env, usando variables de entorno existentes")
	}

	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	sslMode := os.Getenv("DB_SSLMODE")

	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPass, dbHost, dbPort, dbName, sslMode)

	ctx := context.Background()
	pool, err := pgxpool.New(ctx, connStr)
	if err != nil {
		log.Fatalf("❌ Error al conectar con la BD: %v", err)
	}
	defer pool.Close()

	// Test simple
	var now time.Time
	if err := pool.QueryRow(ctx, "SELECT NOW()").Scan(&now); err != nil {
		log.Fatalf("❌ Query falló: %v", err)
	}

	fmt.Println("✅ Conectado a PostgreSQL, hora actual:", now)
}
