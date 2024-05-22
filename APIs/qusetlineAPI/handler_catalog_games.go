package main

import (
	"fmt"

	"github.com/Jett65/questline/APIs/qusetlineAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateCatalogGame(c *fiber.Ctx) error {

	game := new(CatalogGame)

	err := c.BodyParser(game)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to parse body::: %e", err))
	}

	catalogGame, err := apiCfg.DB.CreateCatalogGame(c.Context(), database.CreateCatalogGameParams{
		ID:          uuid.New(),
		Name:        game.Name,
		Description: ToNullString(game.Description),
		Imageurl:    ToNullString(game.ImageURL),
	})
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
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Failed to parse body::: %e", err))
	}

	convCatalogGame, err := catalogGameToDatabeseCatalogGame(game)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to convert body into DB Format: %e", err))
	}

	upGame, err := apiCfg.DB.UpdateCatalogGame(c.Context(), database.UpdateCatalogGameParams{
		ID:          game_id,
		Name:        convCatalogGame.Name,
		Description: convCatalogGame.Description,
		Imageurl:    convCatalogGame.Imageurl,
	})

	payload, err := databaseCatalogGameToCatalogGame(upGame)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to convert catalog game: %e", err))
	}

	return c.JSON(payload)

}

func (apiCfg *apiconfig) handlerGetCatalogGameById(c *fiber.Ctx) error {
	gameId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	catalogGame, err := apiCfg.DB.GetCatalogGameById(c.Context(), gameId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not find game with that id:::: %e", err))
	}

	payload, err := databaseCatalogGameToCatalogGame(catalogGame)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not convert game:::: %e", err))
	}

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerDeleteCatalogGameById(c *fiber.Ctx) error {
	gameId, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Not a valid ID:::: %e", err))
	}

	err = apiCfg.DB.DeleteCatalogGame(c.Context(), gameId)
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("::::Could not delete game with that id:::: %e", err))
	}

	return nil
}
