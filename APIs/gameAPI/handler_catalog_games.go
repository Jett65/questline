package main

import (
	"fmt"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateCatalogGame(c *fiber.Ctx) error {
	
	game := new(CatalogGame)

	err := c.BodyParser(game)
	if err != nil {
		//TODO: Add error message to this error
		return err
	}
 
    convCatalogGame, err := catalogGameToDatabeseCatalogGame(game)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to parse body: %e", err))
    }

	catalogGame, err := apiCfg.DB.CreateCatalogGame(c.Context(), database.CreateCatalogGameParams(convCatalogGame))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to create catalog game: %e", err))
	}

	payload, err := databaseCatalogGameToCatalogGame(catalogGame)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to parse data: %e", err))
	}

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerGetAllCatalogGames(c *fiber.Ctx) error {
	catalogGames, err := apiCfg.DB.GetAllCatalogGames(c.Context())
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to fetch catalog games: %e", err))
	}

	payload, err := databaseCatalogGamesToCatalogGames(catalogGames)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to convert catalog games: %e", err))
	}

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerUpdateCatalogGame(c *fiber.Ctx) error {
	game_id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to parse id: %v", err))
	}

	game := new(CatalogGame)

	err = c.BodyParser(game)
	if err != nil {
		//TODO: Add error message to this error
		return err
	}
 
    convCatalogGame, err := catalogGameToDatabeseCatalogGame(game)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to parse body: %e", err))
    }

	upGame, err := apiCfg.DB.UpdateCatalogGame(c.Context(), database.UpdateCatalogGameParams{
        ID: game_id,
        Name: convCatalogGame.Name,
        Description: convCatalogGame.Description,
        Imageurl: convCatalogGame.Imageurl,
    })

    payload, err := databaseCatalogGameToCatalogGame(upGame)
    if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to convert catalog game: %e", err))
    }

    return c.JSON(payload)
}
