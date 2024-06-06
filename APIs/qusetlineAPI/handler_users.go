package main

import (
	"fmt"
	"time"

	"github.com/Jett65/questline/APIs/qusetlineAPI/internal/database"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
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
		ID:       uuid.New(),
		Username: user.Username,
		Email:    user.Email,
		Password: string(password),
	})

	payload, err := databaseUserToUser(newUser)
	if err != nil {
		return fiber.NewError(404, fmt.Sprintf("::::Failed to create user::::"))
	}
 
    jwt, err := genJWT(secret_key, jwt.MapClaims{
        "iss": "gameAPI", 
        "sub": newUser.Username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    cookie := fiber.Cookie {   
        Name: "key",
        Value: jwt,
        Expires: time.Now().Add(time.Hour * 24),
    } 

    c.Cookie(&cookie)

	return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerlogin(c *fiber.Ctx) error {
    user := new(User)
    
    err := c.BodyParser(&user)
    if err != nil {
        return fiber.NewError(400, "::::Couldn't parse body::::")
    }

	dbUser, err := apiCfg.DB.GetUserByUsername(c.Context(), user.Username)
	if err != nil {
		return fiber.NewError(400, fmt.Sprintf("::::Couldn't find user with that username:::: %e", err))
	}

	parsedUser, err := databaseUserToUser(dbUser)
	if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Failed to parse user::::"))
	}

	err = bcrypt.CompareHashAndPassword([]byte(parsedUser.Password), []byte(user.Password))
    if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Incorrect username or password::::"))
    }

    // Temporary claims discuss with Kden
    jwt, err := genJWT(secret_key, jwt.MapClaims{
        "iss": "gameAPI", 
        "sub": parsedUser.Username, 
        "exp": time.Now().Add(time.Hour * 24).Unix(),
    })

    if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Failed to authenticate user::::"))
    }

    // Figure out the cookie stuff
    cookie := fiber.Cookie {
        Name: "key",
        Value: jwt,
        SameSite: "http://localhost:5173/",
        Expires: time.Now().Add(time.Hour * 24),
    }

    c.Cookie(&cookie)

    payload := User{
        ID: dbUser.ID,
        Username: dbUser.Username,
        Email: dbUser.Email,
        Password: "",
    } 

    return c.JSON(payload)
}

func (apiCfg *apiconfig) handlerDeleteUser(c *fiber.Ctx) error {
    userID, err := uuid.Parse(c. Params("id"))
    if err != nil {
        return fiber.NewError(400, fmt.Sprintf("::::Not a valid user ID::::"))
    }

    err = apiCfg.DB.DeleteUser(c.Context(), userID)
    if err != nil {    
        return fiber.NewError(400, fmt.Sprintf("::::Could not delete User with ID of: %v::::", userID))
    }

    return nil
}

// This handler is a temporary handler to check if the token stuff works
func (apiCfg *apiconfig) isLoggedIn (c *fiber.Ctx) error {
    return c.JSON("The user is logged in")
}
