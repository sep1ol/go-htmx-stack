package http

import (
	"net/http"

	"github.com/gofiber/template/html/v2"
	"github.com/sep1ol/new-stack/pkg/config"
)

func InitHTMLEngine(env config.EnvVars) *html.Engine {
	// Create html engine
	engine := html.NewFileSystem(http.Dir("./html"), ".gohtml")

	// Engine configs
	if env.GO_ENV == "development" {
		engine.Reload(true)
		engine.Debug(true)
	}

	return engine
}
