package utilities

import "github.com/labstack/echo/v4"

// Exception is a custom defined struct containing the details
// relevant to an error in the application
type Exception struct {
	StatusCode int
	Error      string
	Message    string
}

// ThrowException is used to return a JSON error message to the client
func ThrowException(context echo.Context, exception *Exception) error {
	return context.JSON(exception.StatusCode, map[string]string{"error": exception.Error, "message": exception.Message})
}
