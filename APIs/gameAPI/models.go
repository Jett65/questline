package main

import (
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

func databaseBatalogGamesToCatalogGames(dbCatalogGames []database.CatalogGame) ([]CatalogGame, error) {
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
