package main

import (
	"database/sql"
	"fmt"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateCatalogGame(c *fiber.Ctx) error {
	// Used to convert description and imageUrl from
	// string to sql.NullString
	// an sql.NullString is what sqlc uses to handle nullable values
	// for more info here is the link https://pkg.go.dev/database/sql#NullString
	var isNullDes sql.NullString
	var isNullImg sql.NullString

	game := new(CatalogGame)

	err := c.BodyParser(game)
	if err != nil {
		return err
	}

	// converts strings to sqlNullStrings
	if game.Description == "" {
		isNullDes.String = game.Description
		isNullDes.Valid = false
	} else {
		isNullDes.String = game.Description
		isNullDes.Valid = true
	}

	if game.Description == "" {
		isNullImg.String = game.Description
		isNullImg.Valid = false
	} else {
		isNullImg.String = game.Description
		isNullImg.Valid = true
	}

	catalogGame, err := apiCfg.DB.CreateCatalogGame(c.Context(), database.CreateCatalogGameParams{
		ID:          uuid.New(),
		Name:        game.Name,
		Description: isNullDes,
		Imageurl:    isNullImg,
	})
	if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf("failed to create catalog game: %e", err))
	}

	payload, err := databaseCatalogGameToCatalogGame(catalogGame)
	if err != nil {
        return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintf( "failed to parse data: %e", err))
	}

	return c.JSON(payload)
}

// func (apiCfg *apiconfig) handlerCreateCatalogGame(c *fiber.Ctx) error {
//     return c.SendString("it worked")
// }
