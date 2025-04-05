// Package http provides the HTTP Controller implementation for the application.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// PingController is a controller that handles ping requests.
type PingController struct {
}

// NewPingController creates a new PingController.
func NewPingController() *PingController {
	return &PingController{}
}

// RouteInfos returns the route information for the PingController.
// This function is used to register the routes with the Gin engine.
func (controller *PingController) RouteInfos() gin.RoutesInfo {
	return gin.RoutesInfo{
		gin.RouteInfo{
			Method:      "GET",
			Path:        "/api/v1/ping",
			Handler:     "PingController.Ping",
			HandlerFunc: controller.GETPing,
		},
	}
}

// GETPing handles GET requests to the /api/v1/ping endpoint.
func (controller *PingController) GETPing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
