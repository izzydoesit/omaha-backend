// cmd/server/main.go
package main

import (
	"log"
	"os"
	"github.com/gofiber/fiber/v2"
	"github.com/izzydoesit/omaha-backend/internals/models"
	"github.com/izzydoesit/omaha-backend/internal/handlers"
	"github.com/izzydoesit/omaha-backend/internals/storage"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	} 
	dsn := os.Getenv("DATABASE_DSN")
	if dsn == "" {
		log.Fatal("DATABASE_DSN is required in env")
	}
	
	db, err := storage.ConnectDB(dsn)
	if err != nil {
		log.Fatalf("Failed to connect to db: %v", err)
	}
	// Run migrations after connecting to DB
	if err := db.AutoMigrage(&models.Hand{}); err != nil {
		log.Fatalf("DB migration failed: %v", err)
	}

	app := fiber.New()

	handsHandler := &handlers.HandsHandler{DB: db}
	app.Post("/api/hands", handsHandler.CreateHand)
	app.Get("/api/hands", handsHandler.ListHands)
	// app.Get("/api/stats/session", handsHandler.GetSessionStats)
	// app.Get("/api/stats/history", handsHandler.GetHistoryStats)

	log.Println("Server started on port " + os.Getenv("PORT"))
	log.Fatal(app.Listen(":" + os.Getenv("PORT")))
}

// 4 cards = 8 bytes 
// 100bytees per hand
// 10 hands = 1k
// 10000 = 1mb
