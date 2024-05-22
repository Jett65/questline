package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/Jett65/questline/APIs/qusetlineAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
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

	apiCfg := apiconfig{
		DB: database.New(conn),
	}

	app := fiber.New()

	api := app.Group("/api")
	v1 := api.Group("/v1")

    v1.Use(cors.New((cors.Config{
        AllowOrigins: "http://localhost:5173",
    })))

	v1.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("It's Up")
	})

	// Catalog Game Routes
	v1.Post("/catalog_game", apiCfg.handlerCreateCatalogGame)
	v1.Get("/catalog_game", apiCfg.handlerGetAllCatalogGames)
	v1.Get("/catalog_game/:id", apiCfg.handlerGetCatalogGameById)
	v1.Put("/catalog_game/:id", apiCfg.handlerUpdateCatalogGame)
	v1.Delete("/catalog_game/:id", apiCfg.handlerDeleteCatalogGameById)

	// Tasks Routes
	v1.Post("/tasks", apiCfg.handlerCreateTask)
	v1.Get("/tasks", apiCfg.handlerGetAllTasks)
	v1.Get("/tasks/:id", apiCfg.handlerGetTaskById)
	v1.Put("/tasks/:id", apiCfg.handlerUpdateTaskById)
	v1.Delete("/tasks/:id", apiCfg.handlerDeleteTaskById)
	v1.Get("/:gameid/tasks", apiCfg.handlerGetTasksByCatalogGameId)

    // User Routes
    v1.Post("/create_account", apiCfg.handlerCreateUser)
    v1.Post("/login", apiCfg.handlerlogin)
    v1.Get("/is_logged_in", isAuth, apiCfg.isLoggedIn)
    v1.Delete("/users/:id", apiCfg.handlerDeleteUser)

	log.Fatal(app.Listen(":3006"))
}
