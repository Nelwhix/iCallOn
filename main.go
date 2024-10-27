package main

import (
	"fmt"
	"github.com/Nelwhix/iCallOn/handlers"
	"github.com/Nelwhix/iCallOn/pkg"
	"github.com/Nelwhix/iCallOn/pkg/middlewares"
	"github.com/Nelwhix/iCallOn/pkg/models"
	"github.com/go-playground/validator/v10"
	gHandlers "github.com/gorilla/handlers"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

const (
	ServerPort = ":8080"
)

var validate *validator.Validate

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fileName := filepath.Join("logs", "app_logs.txt")
	f, err := os.OpenFile(fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	logger, err := pkg.CreateNewLogger(f)
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
	}

	validate = validator.New(validator.WithRequiredStructEnabled())

	conn, err := pkg.CreateDbConn()
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	model := models.Model{
		Conn: conn,
	}

	handler := handlers.Handler{
		Model:     model,
		Logger:    logger,
		Validator: validate,
	}

	m := middlewares.AuthMiddleware{
		Model: model,
	}

	r := http.NewServeMux()

	// Guest Routes
	r.HandleFunc("POST /api/auth/signup", handler.SignUp)
	r.HandleFunc("POST /api/auth/login", handler.Login)

	// Auth routes
	r.Handle("GET /api/me", m.Register(handler.Me))
	r.Handle("POST /api/games", m.Register(handler.CreateNewGame))

	fmt.Printf("iCallOn started at http://localhost:%s\n", ServerPort)

	err = http.ListenAndServe(ServerPort, gHandlers.CombinedLoggingHandler(os.Stdout, middlewares.ContentTypeMiddleware(r)))

	if err != nil {
		log.Printf("failed to run the server: %v", err)
	}
}
