package controllers

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
	"time"
	"user_management/models"
	"user_management/utils"
)

// UserController handles user-related requests
type UserController struct {
	Repo models.UserRepository
}

// CreateUser handles POST /users request to create a new user.
func (uc *UserController) CreateUser(c echo.Context) error {
	start := time.Now()
	defer utils.LogRequest(start, "/users")

	var user models.User
	if err := c.Bind(&user); err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	if exists, err := uc.Repo.EmailExists(user.Email); err != nil {
		utils.LogError(err.Error())
		return utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to check email")
	} else if exists {
		return utils.SendErrorResponse(c, http.StatusConflict, "Email already exists")
	}

	id, err := uc.Repo.CreateUser(user)
	if err != nil {
		utils.LogError(err.Error())
		return utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to create user")
	}

	return utils.SendSuccessResponse(c, http.StatusCreated, "User created", map[string]int{"id": id})
}

// GetUser handles GET /users/{id} request to retrieve user details.
func (uc *UserController) GetUser(c echo.Context) error {
	start := time.Now()
	defer utils.LogRequest(start, "/users/"+c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	user, err := uc.Repo.GetUser(id)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
	}

	return utils.SendSuccessResponse(c, http.StatusOK, "User retrieved", user)
}

// UpdateUser handles PUT /users/{id} request to update user details.
func (uc *UserController) UpdateUser(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid user ID")
	}

	var user models.User
	if err := c.Bind(&user); err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid input")
	}

	user.ID = id

	exists, err := uc.Repo.UserExists(id)
	if err != nil || !exists {
		return utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
	}

	if err := uc.Repo.UpdateUser(id, user); err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to update user")
	}

	//return c.JSON(http.StatusOK, user)
	return utils.SendSuccessResponse(c, http.StatusOK, "User updated", nil)

}

// DeleteUser handles DELETE /users/{id} request to delete a user.
func (uc *UserController) DeleteUser(c echo.Context) error {
	start := time.Now()
	defer utils.LogRequest(start, "/users/"+c.Param("id"))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusBadRequest, "Invalid ID")
	}

	exists, err := uc.Repo.UserExists(id)
	if err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to check user existence")
	} else if !exists {
		return utils.SendErrorResponse(c, http.StatusNotFound, "User not found")
	}

	if err := uc.Repo.DeleteUser(id); err != nil {
		return utils.SendErrorResponse(c, http.StatusInternalServerError, "Failed to delete user")
	}

	return utils.SendSuccessResponse(c, http.StatusOK, "User deleted", nil)
}
