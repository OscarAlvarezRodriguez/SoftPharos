package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"softpharos/cmd/app"
	"softpharos/internal/infra/databases"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("⚠️  No se pudo cargar .env, usando variables de entorno existentes")
	}

	// Crear e inicializar la instancia única de base de datos (Singleton)
	dbClient, err := databases.NewClientFromEnv()
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}
	databases.InitializeDatabase(dbClient)
	defer databases.CloseInstance()

	// Verificar conexión
	if err := databases.GetInstance().Ping(); err != nil {
		log.Fatalf("❌ Error al hacer ping a la base de datos: %v", err)
	}

	// Configurar modo de Gin
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// Crear router
	router := gin.Default()

	// Configurar CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Mapear rutas (cada dominio registra sus propias rutas)
	app.MapUrls(router)

	// Obtener puerto
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Iniciar servidor
	if err := router.Run(":" + port); err != nil {
		log.Fatalf("❌ Error al iniciar el servidor: %v", err)
	}
}
