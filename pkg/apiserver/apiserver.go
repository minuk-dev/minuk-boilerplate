// Package apiserver provides the API server implementation for the application.
// It uses the Uber Fx framework for dependency injection and lifecycle management.
// It also uses the Gin framework for HTTP handling and routing.
package apiserver

import (
	"net/http"

	"go.uber.org/fx"

	httpadapter "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/adapter/in/http"
	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/adapter/out/sqlite"
	domainin "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/in"
	domainout "github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/port/out"
	"github.com/minuk-dev/minuk-boilerplate/pkg/apiserver/domain/service"
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
	DB   string
}

// New creates a new APIServer instance with the provided settings.
func New(settings Settings) *APIServer {
	app := fx.New(
		// base
		fx.Provide(
			NewHTTPServer,
			NewLogger,
			fx.Annotate(
				NewEngine,
				fx.ParamTags(`group:"controller"`),
			),
		),
		// controllers
		fx.Provide(
			fx.Annotate(httpadapter.NewPingController, fx.As(new(Controller)), fx.ResultTags(`group:"controller"`)),
		),
		// services
		fx.Provide(
			fx.Annotate(service.NewHistoryService, fx.As(new(domainin.HistoryUsecase))),
		),
		// repositories
		fx.Provide(
			fx.Annotate(sqlite.NewHistorySqliteAdapter, fx.As(new(domainout.HistoryRepository))),
		),
		// database
		fx.Provide(
			NewDB,
		),

		// settings
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
