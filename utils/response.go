package utils

import "github.com/labstack/echo/v4"

// SendSuccessResponse sends a success response with a JSON body
func SendSuccessResponse(c echo.Context, statusCode int, message string, data interface{}) error {
	response := map[string]interface{}{
		"status":  statusCode,
		"message": message,
		"data":    data,
	}
	return c.JSON(statusCode, response)
}

// SendErrorResponse sends an error response with a JSON body
func SendErrorResponse(c echo.Context, statusCode int, message string) error {
	response := map[string]interface{}{
		"status": statusCode,
		"error":  message,
	}
	return c.JSON(statusCode, response)
}
