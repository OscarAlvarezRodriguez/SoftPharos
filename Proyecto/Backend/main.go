package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"softpharos/cmd/app"
	"softpharos/cmd/buildingAPI"
	"softpharos/internal/infra/databases"
)

func main() {
	// Cargar variables de entorno
	if err := godotenv.Load("../../.env"); err != nil {
		log.Println("⚠️  No se pudo cargar .env, usando variables de entorno existentes")
	}

	// Crear cliente de base de datos
	dbClient, err := databases.NewClientFromEnv()
	if err != nil {
		log.Fatalf("❌ Error al conectar con la base de datos: %v", err)
	}
	defer dbClient.Close()

	// Verificar conexión
	if err := dbClient.Ping(); err != nil {
		log.Fatalf("❌ Error al hacer ping a la base de datos: %v", err)
	}

	// Construir dependencias (inyección de dependencias)
	projectControllers := buildingAPI.BuildProjectController(dbClient)
	roleControllers := buildingAPI.BuildRoleController(dbClient)

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

	// Mapear rutas
	app.MapUrls(router, projectControllers, roleControllers)

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
