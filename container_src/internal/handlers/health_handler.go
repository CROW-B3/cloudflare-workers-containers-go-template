package handlers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"server/config"
	"server/pkg/database"
	"server/pkg/response"
)

type HealthHandler struct {
	config *config.Config
}

func NewHealthHandler(cfg *config.Config) *HealthHandler {
	return &HealthHandler{config: cfg}
}

// HealthCheck godoc
// @Summary Health check endpoint
// @Description Returns the health status of the application
// @Tags health
// @Produce json
// @Success 200 {object} response.Response
// @Router /health [get]
func (h *HealthHandler) HealthCheck(c *gin.Context) {
	response.Success(c, http.StatusOK, "Service is healthy", gin.H{
		"status":    "ok",
		"timestamp": time.Now().Format(time.RFC3339),
		"service":   h.config.App.Name,
		"version":   h.config.App.Version,
	})
}

// ReadinessCheck godoc
// @Summary Readiness check endpoint
// @Description Returns the readiness status of the application including database connectivity
// @Tags health
// @Produce json
// @Success 200 {object} response.Response
// @Router /ready [get]
func (h *HealthHandler) ReadinessCheck(c *gin.Context) {
	// Check database connection
	dbStatus := "ok"
	sqlDB, err := database.DB.DB()
	if err != nil || sqlDB.Ping() != nil {
		dbStatus = "unhealthy"
	}

	status := "ready"
	statusCode := http.StatusOK
	if dbStatus != "ok" {
		status = "not ready"
		statusCode = http.StatusServiceUnavailable
	}

	response.Success(c, statusCode, status, gin.H{
		"database":  dbStatus,
		"timestamp": time.Now().Format(time.RFC3339),
	})
}

// LivenessCheck godoc
// @Summary Liveness check endpoint
// @Description Returns the liveness status of the application
// @Tags health
// @Produce json
// @Success 200 {object} response.Response
// @Router /live [get]
func (h *HealthHandler) LivenessCheck(c *gin.Context) {
	response.Success(c, http.StatusOK, "Service is alive", gin.H{
		"status":      "alive",
		"timestamp":   time.Now().Format(time.RFC3339),
		"instance_id": os.Getenv("CLOUDFLARE_DURABLE_OBJECT_ID"),
	})
}
