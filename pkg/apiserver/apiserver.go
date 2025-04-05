// Package apiserver provides the API server implementation for the application.
// It uses the Uber Fx framework for dependency injection and lifecycle management.
// It also uses the Gin framework for HTTP handling and routing.
package apiserver

import (
	"net/http"

	"go.uber.org/fx"

	httpadapter "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/adapter/in/http"
)

// APIServer is a struct that represents the API server.
// It embeds the fx.App struct from the Uber Fx framework.
type APIServer struct {
	*fx.App
}

// Settings is a struct that holds the configuration settings for the API server.
// It contains the address to bind the server to.
type Settings struct {
	Addr string
}

// New creates a new APIServer instance with the provided settings.
func New(settings Settings) *APIServer {
	app := fx.New(
		fx.Provide(
			NewHTTPServer,
			NewLogger,
			fx.Annotate(
				NewEngine,
				fx.ParamTags(`group:"controller"`),
			),
		),
		fx.Provide(
			fx.Annotate(httpadapter.NewPingController, fx.As(new(Controller)), fx.ResultTags(`group:"controller"`)),
		),
		fx.Provide(func() *Settings {
			return &settings
		}),

		// to Triger the lifecycle of the server
		fx.Invoke(func(*http.Server) {}),
	)

	return &APIServer{
		App: app,
	}
}
