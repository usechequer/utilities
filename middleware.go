package utilities

import (
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

// Token is a custom defined struct that represents relevant information
// decoded from the JWT Token
type Token struct {
	Subject string
	Issuer  string
}

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

		subject, _ := decodedToken.Claims.GetSubject()
		issuer, _ := decodedToken.Claims.GetIssuer()

		contextToken := Token{Subject: subject, Issuer: issuer}
		context.Set("token", contextToken)

		return next(context)
	}
}
