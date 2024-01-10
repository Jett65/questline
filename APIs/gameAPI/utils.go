package main

import (
	"database/sql"
)

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

// func handleErr(c *fiber.Ctx, message string, err error) error {
// 	if err == nil {
// 		return fiber.NewError(fiber.StatusBadRequest, fmt.Sprintln("::::"+message+"::::", "%e"), err.Error())
// 	}
// 	return nil
// }
