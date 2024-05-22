package main

import (
	"github.com/Jett65/questline/APIs/qusetlineAPI/internal/database"
	"github.com/google/uuid"
)

// CatalogGames
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
	return database.CatalogGame{
		ID:          catalogGmae.ID,
		Name:        catalogGmae.Name,
		Description: ToNullString(catalogGmae.Description),
		Imageurl:    ToNullString(catalogGmae.ImageURL),
	}, nil
}

// Tasks
type Task struct {
	ID          uuid.UUID `json:"id"`
	Description string    `json:"description"`
	Game_id     uuid.UUID `json:"game_id"`
}

func databaseTaskToTask(dbTask database.Task) (Task, error) {
	return Task{
		ID:          dbTask.ID,
		Description: dbTask.Description.String,
		Game_id:     dbTask.GameID,
	}, nil
}

func databasetasksToTasks(dbTasks []database.Task) ([]Task, error) {
	tasks := []Task{}
	for _, task := range dbTasks {
		dbGame, err := databaseTaskToTask(task)
		if err != nil {
			return tasks, err
		}

		tasks = append(tasks, dbGame)
	}

	return tasks, nil
}

func taskToDatabaseTask(task Task) (database.Task, error) {
	return database.Task{
		ID:          task.ID,
		Description: ToNullString(task.Description),
		GameID:      task.Game_id,
	}, nil
}

// Users
type User struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
}

//TODO: Create email validation

func databaseUserToUser(dbUser database.User) (User, error) {
	return User{
		ID:       dbUser.ID,
		Username: dbUser.Username,
		Email:    dbUser.Email,
		Password: dbUser.Password,
	}, nil
}

func databaseUsersToUsers(dbUsers []database.User) ([]User, error) {
	users := []User{}

	for _, user := range dbUsers {
		dbUser, err := databaseUserToUser(user)
		if err != nil {
			return users, err
		}

		users = append(users, dbUser)
	}

	return users, nil
}
