package router

import (
	"html/template"
	"interface/internal/config"
	"interface/internal/http-server/handlers"
	asciiart "interface/internal/processor/ascii-art"
	"log/slog"
	"net/http"
)

func addRoutes(mux *http.ServeMux, logger *slog.Logger, cfg *config.Config, artProc *asciiart.ASCIIArtProcessor, tmplts map[string]*template.Template) {

	// Page handlers
	mux.Handle("GET /{$}", handlers.HandleIndex(logger, tmplts)) // handle must return http.Handler
	mux.Handle("GET /web/", handlers.HandleWebAssets(logger, cfg))
	mux.Handle("GET /", handlers.HandleNotFound(logger, tmplts))

	// Action handlers
	mux.Handle("POST /decoder", handlers.HandleDecoder(logger, artProc, tmplts))
	mux.Handle("POST /encoder", handlers.HandleEncoder(logger, artProc, tmplts))
	mux.Handle("POST /randomizer", handlers.HandleRandomArt(logger, cfg, tmplts))
}

// TODO: maybe wrap func newMiddleware(logger Logger) func(h http.Handler) http.Handler
