package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/GustavoMS97/go-notes-api/docs"
	"github.com/GustavoMS97/go-notes-api/internal/app"
	"github.com/joho/godotenv"
)

// @title Go Notes API
// @version 1.0
// @description API for notes with JWT authentication
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @BasePath /api
func main() {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		log.Println("Exiting with Ctrl+C")
		os.Exit(0)
	}()

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	// Set the host dynamically from .env or fallback
	docs.SwaggerInfo.Host = os.Getenv("SWAGGER_HOST")
	if docs.SwaggerInfo.Host == "" {
		docs.SwaggerInfo.Host = "localhost:4000"
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	env := os.Getenv("ENV")
	log.Printf("Running in %s mode...\n", env)

	fiberApp := app.InitApp()

	log.Fatal(fiberApp.Listen(":" + port))
}
