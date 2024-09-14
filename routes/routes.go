package routes

import (
	"github.com/labstack/echo/v4"
	"user_management/controllers"
)

// SetupRoutes sets up the API routes for user management.
func SetupRoutes(e *echo.Echo, controllers *controllers.UserController) {
	e.POST("/users", controllers.CreateUser)
	e.GET("/users/:id", controllers.GetUser)
	e.PUT("/users/:id", controllers.UpdateUser)
	e.DELETE("/users/:id", controllers.DeleteUser)
}
