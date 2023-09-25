package http

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/sep1ol/new-stack/config"
	"github.com/sep1ol/new-stack/handlers"
	"github.com/sep1ol/new-stack/pkg/database"
)

func ServeHTTP(env config.EnvVars) (func(), error) {
	app, cleanup, err := buildServer(env)
	if err != nil {
		return nil, err
	}

	// Start server
	go func() {
		app.Listen("0.0.0.0:" + env.PORT)
	}()

	// Return: close server and db func
	return func() {
		cleanup()
		app.Shutdown()
	}, nil
}

func buildServer(env config.EnvVars) (*fiber.App, func(), error) {
	// Connect to database
	db, err := database.ConnectPostgres(env.DB_URL)
	if err != nil {
		return nil, nil, err
	}

	// Create Fiber App with html engine
	HTMLEngine := html.New("./html", ".gohtml")
	app := fiber.New(fiber.Config{
		Views: HTMLEngine,
	})

	// Serve static files
	app.Static("/", "./public")

	// Routes & Middlewares
	handlers.RegisterAPIHandlers(app, db)
	handlers.RegisterClientHandlers(app, db)

	return app, func() {
		database.ClosePostgres(db)
	}, nil
}
