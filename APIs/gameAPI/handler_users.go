package main

import (
	"fmt"

	"github.com/Jett65/questline/APIs/gameAPI/internal/database"
	"github.com/gofiber/fiber/v2"
    "golang.org/x/crypto/bcrypt"
	"github.com/google/uuid"
)

func (apiCfg *apiconfig) handlerCreateUser(c *fiber.Ctx) error {
    user := new(User)

    err := c.BodyParser(&user)
    if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Failed to parse body:::: %e", err))
    }

    password, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
    if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Failed to create user::::"))
    }

    newUser, err := apiCfg.DB.CreateUser(c.Context(), database.CreateUserParams{
        ID: uuid.New(), 
        Username: user.Username,
        Email: user.Email,
        Password: string(password),
    })

    payload, err := databaseUserToUser(newUser)
    if err != nil {
        return fiber.NewError(404, fmt.Sprintf("::::Failed to create user::::"))
    }

    return c.JSON(payload)
}

// func (apiCfg *apiconfig) handlerlogin(c *fiber.Ctx) error {
//     
// }
