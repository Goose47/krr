// Package app defines application model.
package app

import (
	"context"
	serverapp "krr/internal/app/server"
	"krr/internal/controllers"
	"krr/internal/server"
	"krr/internal/services"
	"log/slog"
)

type clearer interface {
	Clear(ctx context.Context)
}

// App represents application.
type App struct {
	Server *serverapp.Server
	Cache  clearer
}

// New creates all dependencies for App and returns new App instance.
func New(
	log *slog.Logger,
	env string,
	port int,
) *App {
	testService := services.NewTestulatorService(log)

	testCon := controllers.NewTestController(log, testService)

	router := server.NewRouter(log, env, testCon)
	serverApp := serverapp.New(log, port, router)

	return &App{
		Server: serverApp,
	}
}
