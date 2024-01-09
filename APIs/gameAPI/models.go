package main

import (
	"database/sql"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/google/uuid"
)

type CatalogGame struct {
	ID          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	ImageURL    string    `json:"image_url"`
}

func databaseCatalogGameToCatalogGame(dbCatalogGame database.CatalogGame) (CatalogGame, error) {
	return CatalogGame{
		ID:          dbCatalogGame.ID,
		Name:        dbCatalogGame.Name,
		Description: dbCatalogGame.Description.String,
		ImageURL:    dbCatalogGame.Imageurl.String,
	}, nil
}

func databaseCatalogGamesToCatalogGames(dbCatalogGames []database.CatalogGame) ([]CatalogGame, error) {
	games := []CatalogGame{}
	for _, game := range dbCatalogGames {
		dbGame, err := databaseCatalogGameToCatalogGame(game)
		if err != nil {
			return games, err
		}

		games = append(games, dbGame)
	}

	return games, nil
}

func catalogGameToDatabeseCatalogGame(catalogGmae *CatalogGame) (database.CatalogGame, error) {
	var isNullDes sql.NullString
	var isNullImg sql.NullString

	if catalogGmae.Description == "" {
		isNullDes.String = catalogGmae.Description
		isNullDes.Valid = false
	} else {
		isNullDes.String = catalogGmae.Description
		isNullDes.Valid = true
	}

	if catalogGmae.ImageURL == "" {
		isNullImg.String = catalogGmae.ImageURL
		isNullImg.Valid = false
	} else {
		isNullImg.String = catalogGmae.ImageURL
		isNullImg.Valid = true
	}

	return database.CatalogGame{
		ID:          catalogGmae.ID,
		Name:        catalogGmae.Name,
		Description: isNullDes,
		Imageurl:    isNullImg,
	}, nil
}
