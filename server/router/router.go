package router

import (
	"html/template"
	"interface/internal/config"
	"interface/internal/http-server/helpers"
	"interface/internal/http-server/middleware/logger"
	"interface/internal/http-server/middleware/recoverer"
	reqid "interface/internal/http-server/middleware/requestid"
	asciiart "interface/internal/processor/ascii-art"
	"log/slog"
	"net/http"
)

func NewRouter(log *slog.Logger, cfg *config.Config, artProc *asciiart.ASCIIArtProcessor, tmplts map[string]*template.Template) http.Handler {
	mux := http.NewServeMux()

	addRoutes(
		mux,
		log,
		cfg,
		artProc,
		tmplts,
	)

	reqID := reqid.NewReqIDMiddleware(log)
	logMw := logger.NewLoggingMiddleware(log)
	recoverMw := recoverer.NewRecoveringMiddleware(log)
	handler := helpers.Chain(mux, reqID, logMw, recoverMw)
	return handler
}

// func newMiddleware(log *slog.Logger) func(h http.Handler) http.Handler
