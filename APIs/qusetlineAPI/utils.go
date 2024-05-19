package main

import (
	"database/sql"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

var secret_key = []byte(os.Getenv("SECRET_KEY"))

func genJWT(key []byte, climes jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, climes)

	tokenString, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// middlewhere to authorize the user
func isAuth(c *fiber.Ctx) error {
    tokenString := c.Cookies("key")

    parseToken, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
        return secret_key, nil
    })
    if err != nil {
        return fiber.NewError(400, "::::Failed To parse Token::::")
    }

    if !parseToken.Valid {
        return fiber.NewError(400, "::::Token is invalid::::")
    }

    return c.Next()
}

func ParseNullString(nullString *sql.NullString) string {
	return nullString.String
}

func ToNullString(str string) sql.NullString {
	if str == "" {
		return sql.NullString{
			String: str,
			Valid:  false,
		}
	} else {
		return sql.NullString{
			String: str,
			Valid:  true,
		}
	}
}

