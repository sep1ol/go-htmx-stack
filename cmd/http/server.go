package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sep1ol/new-stack/handlers"
	"github.com/sep1ol/new-stack/pkg/config"
	"github.com/sep1ol/new-stack/pkg/database"
)

func ServeHTTP(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	db, err := database.ConnectPostgres(env.DB_URL)
	if err != nil {
		return nil, nil, err
	}

	engine := InitHTMLEngine(env)
	app := fiber.New(fiber.Config{
		Views: engine,
	})

	app.Static("/", "./public")

	handlers.UseHTMXHandlers(app, db)
	handlers.UseRoutingHandlers(app, db)

	return app, func() {
		database.ClosePostgres(db)
	}, nil
}
