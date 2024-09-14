package controllers_test

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
	"user_management/controllers"
	"user_management/mocks"
	"user_management/models"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

// Helper function to set up a new Echo context
func newContext(method, path, body string) (echo.Context, *httptest.ResponseRecorder) {
	e := echo.New()
	req := httptest.NewRequest(method, path, bytes.NewBufferString(body))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	return c, rec
}

func TestCreateUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	user := models.User{Name: "John Doe", Email: "john@example.com", Age: 30}
	mockRepo.On("EmailExists", user.Email).Return(false, nil)
	mockRepo.On("CreateUser", user).Return(1, nil)

	c, rec := newContext(http.MethodPost, "/users", `{"name":"John Doe","email":"john@example.com","age":30}`)
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.CreateUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusCreated, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestCreateUser_EmailAlreadyExists(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	user := models.User{Name: "John Doe", Email: "john@example.com", Age: 30}
	mockRepo.On("EmailExists", user.Email).Return(true, nil)

	c, rec := newContext(http.MethodPost, "/users", `{"name":"John Doe","email":"john@example.com","age":30}`)
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.CreateUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusConflict, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	user := models.User{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}
	mockRepo.On("GetUser", 1).Return(user, nil)

	c, rec := newContext(http.MethodGet, "/users/1", "")
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.GetUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestGetUser_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	mockRepo.On("GetUser", 1).Return(models.User{}, echo.ErrNotFound)

	c, rec := newContext(http.MethodGet, "/users/1", "")
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.GetUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	user := models.User{ID: 1, Name: "John Doe", Email: "john@example.com", Age: 30}
	mockRepo.On("UserExists", user.ID).Return(true, nil)
	mockRepo.On("UpdateUser", user.ID, user).Return(nil)

	c, rec := newContext(http.MethodPut, "/users/1", `{"name":"John Doe","email":"john@example.com","age":30}`)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.UpdateUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestUpdateUser_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	mockRepo.On("UserExists", 1).Return(false, nil)

	c, rec := newContext(http.MethodPut, "/users/1", `{"name":"John Doe","email":"john@example.com","age":30}`)
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.UpdateUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	mockRepo.On("UserExists", 1).Return(true, nil)
	mockRepo.On("DeleteUser", 1).Return(nil)

	c, rec := newContext(http.MethodDelete, "/users/1", "")
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.DeleteUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, rec.Code)
	mockRepo.AssertExpectations(t)
}

func TestDeleteUser_NotFound(t *testing.T) {
	// Arrange
	mockRepo := new(mocks.MockUserRepository)
	mockRepo.On("UserExists", 1).Return(false, nil)

	c, rec := newContext(http.MethodDelete, "/users/1", "")
	c.SetPath("/users/:id")
	c.SetParamNames("id")
	c.SetParamValues("1")
	uc := &controllers.UserController{Repo: mockRepo}

	// Act
	err := uc.DeleteUser(c)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, http.StatusNotFound, rec.Code)
	mockRepo.AssertExpectations(t)
}
