// Package apiserver provides the HTTP server implementation for the application.
package apiserver

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

const (
	// DefaultReadHeaderTimeout is the default timeout for reading the header of a request.
	// It is set to 30 seconds.
	// This timeout is used to prevent the server from waiting indefinitely for a request header
	// because of security reasons.
	DefaultReadHeaderTimeout = 30 * time.Second
)

// NewHTTPServer creates a new HTTP server instance.
func NewHTTPServer(
	lifecycle fx.Lifecycle,
	engine *gin.Engine,
	logger *slog.Logger,
) *http.Server {
	//exhaustruct:ignore
	srv := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: DefaultReadHeaderTimeout,
		Handler:           engine,
	}

	// srv.ConnContext = opampController.ConnContext
	lifecycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			listener, err := net.Listen("tcp", srv.Addr)
			if err != nil {
				return fmt.Errorf("failed to listen on %s: %w", srv.Addr, err)
			}
			logger.Info("Starting HTTP Server",
				slog.String("addr", srv.Addr))
			go func() {
				err := srv.Serve(listener)
				if err != nil {
					logger.Error("HTTP Server error",
						slog.String("addr", srv.Addr),
						slog.String("error", err.Error()))
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return srv
}

// NewEngine creates a new Gin engine instance.
func NewEngine(controllers []Controller) *gin.Engine {
	engine := gin.New()
	engine.Use(gin.Recovery())
	engine.Use(gin.Logger())

	for _, controller := range controllers {
		for _, route := range controller.RouteInfos() {
			engine.Handle(route.Method, route.Path, route.HandlerFunc)
		}
	}

	return engine
}

// Controller is an interface that defines the methods for a controller.
// It is used to define the contract for all controllers in the application.
// Each controller should implement this interface to be registered with the Gin engine.
type Controller interface {
	RouteInfos() gin.RoutesInfo
}
