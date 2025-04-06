// Package http provides the HTTP Controller implementation for the application.
package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	apiv1 "github.com/minuk-dev/minuk-boilerplate/api/v1"
	domainin "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/in"
)

// PingController is a controller that handles ping requests.
type PingController struct {
	historyUsecase domainin.HistoryUsecase
}

// NewPingController creates a new PingController.
func NewPingController(
	historyUsecase domainin.HistoryUsecase,
) *PingController {
	return &PingController{
		historyUsecase: historyUsecase,
	}
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
func (controller *PingController) GETPing(ginC *gin.Context) {
	err := controller.historyUsecase.Save(
		uuid.New(),
		ginC.Request.Header,
	)
	if err != nil {
		ginC.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})

		return
	}

	response := apiv1.PingResponse{
		Message: "pong",
	}

	ginC.JSON(http.StatusOK, response)
}
