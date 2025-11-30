package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   interface{} `json:"error,omitempty"`
}

// Success sends a successful JSON response
func Success(c *gin.Context, statusCode int, message string, data interface{}) {
	c.JSON(statusCode, Response{
		Success: true,
		Message: message,
		Data:    data,
	})
}

// Error sends an error JSON response
func Error(c *gin.Context, statusCode int, message string, err interface{}) {
	c.JSON(statusCode, Response{
		Success: false,
		Message: message,
		Error:   err,
	})
}

// ValidationError sends a validation error response
func ValidationError(c *gin.Context, errors interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Success: false,
		Message: "Validation failed",
		Error:   errors,
	})
}

// InternalError sends an internal server error response
func InternalError(c *gin.Context, message string) {
	c.JSON(http.StatusInternalServerError, Response{
		Success: false,
		Message: message,
		Error:   "Internal server error",
	})
}

// NotFound sends a not found error response
func NotFound(c *gin.Context, message string) {
	c.JSON(http.StatusNotFound, Response{
		Success: false,
		Message: message,
		Error:   "Resource not found",
	})
}

// Unauthorized sends an unauthorized error response
func Unauthorized(c *gin.Context, message string) {
	c.JSON(http.StatusUnauthorized, Response{
		Success: false,
		Message: message,
		Error:   "Unauthorized",
	})
}
