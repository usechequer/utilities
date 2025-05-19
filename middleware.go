package utilities

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// AuthMiddleware is a middleware that inspects the request's authorization
// header for the relevant token, and if it exists, checks if it is a valid
// token issued by Carbon. If it is not, it throws an error.
func AuthMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(context echo.Context) error {
		Authorization := context.Request().Header.Get("Authorization")

		authHeaderSplits := strings.Split(Authorization, " ")

		if len(authHeaderSplits) != 2 {
			return ThrowException(context, &Exception{StatusCode: http.StatusUnauthorized, Message: "Not authenticated", Error: "AUTH_004"})
		}

		token := authHeaderSplits[1]

		decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("JWT_SECRET_KEY")), nil
		})

		if err != nil || !decodedToken.Valid {
			return ThrowException(context, &Exception{StatusCode: http.StatusUnauthorized, Message: "Not authenticated", Error: "AUTH_004"})
		}

		// userUuid, _ := decodedToken.Claims.GetSubject()

		// var user models.User

		// database := GetDatabaseObject()

		// result := database.Where("uuid = ?", userUuid).First(&user)

		// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		// 	return ThrowException(context, &Exception{StatusCode: http.StatusUnauthorized, Message: "Not authenticated", Error: "AUTH_004"})
		// }

		// context.Set("user", user)

		return next(context)
	}
}
