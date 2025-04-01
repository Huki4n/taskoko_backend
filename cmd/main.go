package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"

	"awesomeProject/ent"
	"awesomeProject/internal/routes"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

const (
	Reset     = "\033[0m"
	BlackBg   = "\033[40m"
	RedBg     = "\033[41m"
	GreenBg   = "\033[42m"
	YellowBg  = "\033[43m"
	BlueBg    = "\033[44m"
	MagentaBg = "\033[45m"
	CyanBg    = "\033[46m"
	WhiteBg   = "\033[47m"
	Bold      = "\033[1m"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	escapedPassword := url.QueryEscape(dbPassword)

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", dbUser, escapedPassword, dbHost, dbPort, dbName)

	client, err := ent.Open("postgres", dsn)
	if err != nil {
		fmt.Printf("failed opening connection to database: %v", err)
	}
	defer func(client *ent.Client) {
		err := client.Close()
		if err != nil {
			fmt.Printf("failed closing connection to database: %v", err)
		}
	}(client)

	// Автоматическая миграция схемы (создание таблицы Message, если её нет)
	if err := client.Schema.Create(context.Background()); err != nil {
		fmt.Printf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: MagentaBg + "[${time_rfc3339}]" + Reset +
			" | " + BlueBg + Bold + " ${method} " + Reset +
			" ${uri} | " + GreenBg + " Code: ${status} " + Reset +
			" | ⏱ ${latency_human} | 🔽 ${bytes_in}B 🔼 ${bytes_out}B\n",
		Output: os.Stdout,
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Set-Cookie"},
	}))

	routes.RegisterRoutes(e, client)

	err = e.Start("localhost:8080")

	if err != nil {
		return
	}
}
