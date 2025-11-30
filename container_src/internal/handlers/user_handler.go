package handlers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"server/internal/models"
	"server/internal/services"
	"server/pkg/response"
	"server/pkg/validator"
)

type UserHandler struct {
	service services.UserService
}

func NewUserHandler(service services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

// CreateUser godoc
// @Summary Create a new user
// @Description Create a new user with the provided details
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users [post]
func (h *UserHandler) CreateUser(c *gin.Context) {
	var req models.CreateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate request
	if validationErrors := validator.ValidateStruct(&req); validationErrors != nil {
		response.ValidationError(c, validationErrors)
		return
	}

	user, err := h.service.CreateUser(&req)
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Failed to create user", err.Error())
		return
	}

	response.Success(c, http.StatusCreated, "User created successfully", user)
}

// GetUser godoc
// @Summary Get user by ID
// @Description Get a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /users/{id} [get]
func (h *UserHandler) GetUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "User not found")
			return
		}
		response.InternalError(c, "Failed to get user")
		return
	}

	response.Success(c, http.StatusOK, "User retrieved successfully", user)
}

// GetUsers godoc
// @Summary Get all users
// @Description Get a paginated list of users
// @Tags users
// @Produce json
// @Param limit query int false "Limit" default(10)
// @Param offset query int false "Offset" default(0)
// @Success 200 {object} response.Response
// @Router /users [get]
func (h *UserHandler) GetUsers(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	users, err := h.service.GetUsers(limit, offset)
	if err != nil {
		response.InternalError(c, "Failed to get users")
		return
	}

	response.Success(c, http.StatusOK, "Users retrieved successfully", gin.H{
		"users":  users,
		"limit":  limit,
		"offset": offset,
		"count":  len(users),
	})
}

// UpdateUser godoc
// @Summary Update user
// @Description Update a user by their ID
// @Tags users
// @Accept json
// @Produce json
// @Param id path string true "User ID"
// @Param user body models.UpdateUserRequest true "User update details"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Router /users/{id} [put]
func (h *UserHandler) UpdateUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	var req models.UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid request body", err.Error())
		return
	}

	// Validate request
	if validationErrors := validator.ValidateStruct(&req); validationErrors != nil {
		response.ValidationError(c, validationErrors)
		return
	}

	user, err := h.service.UpdateUser(id, &req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "User not found")
			return
		}
		response.Error(c, http.StatusBadRequest, "Failed to update user", err.Error())
		return
	}

	response.Success(c, http.StatusOK, "User updated successfully", user)
}

// DeleteUser godoc
// @Summary Delete user
// @Description Delete a user by their ID
// @Tags users
// @Produce json
// @Param id path string true "User ID"
// @Success 200 {object} response.Response
// @Failure 404 {object} response.Response
// @Router /users/{id} [delete]
func (h *UserHandler) DeleteUser(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		response.Error(c, http.StatusBadRequest, "Invalid user ID", err.Error())
		return
	}

	err = h.service.DeleteUser(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.NotFound(c, "User not found")
			return
		}
		response.InternalError(c, "Failed to delete user")
		return
	}

	response.Success(c, http.StatusOK, "User deleted successfully", nil)
}
