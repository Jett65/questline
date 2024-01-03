package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type apiconfig struct {
    DB *database.Queries
}

func main() {    
    godotenv.Load(".env")

    dbURL := os.Getenv("DATABASE_URL")
    if dbURL == "" {
        log.Fatal("DATABASE_URL is not found in the environment")
    }

    conn, err := sql.Open("postgres", dbURL)
    if err != nil {
        log.Fatal("Can't connect to database")
    }

    apiCfg := apiconfig {
        DB: database.New(conn),
    }

    app := fiber.New()

    api := app.Group("/api") 
    v1 := api.Group("/v1")

    v1.Get("/test", func(c *fiber.Ctx) error {
        test, err := apiCfg.DB.GetTest(c.Context())    
        if err != nil { 
            return c.Status(http.StatusBadRequest).SendString(fmt.Sprintf("%v", err))
        }

        return c.Status(http.StatusOK).SendString(test)
    })

    log.Fatal(app.Listen(":3006"))
}
