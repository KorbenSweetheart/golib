package handlers

import (
	"bytes"
	"html/template"
	"interface/internal/http-server/helpers"
	asciiart "interface/internal/processor/ascii-art"
	"log/slog"
	"net/http"
	"strings"
)

func HandleEncoder(log *slog.Logger, artProc *asciiart.ASCIIArtProcessor, tmplts map[string]*template.Template) http.Handler {
	const op = "handlers.encoder.handle.Encoder"

	log = log.With(
		slog.String("op", op),
	)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		var (
			sb                 strings.Builder
			HTTPResponseHeader string
			result             string
			statusCode         int
		)

		indexTmpl, ok := tmplts["index"]
		if !ok {
			log.Error("template not found in map", "name", "index")
			http.Error(w, "Internal Configuration Error", http.StatusInternalServerError)
			return
		}

		input := r.FormValue("textinput")
		if input == "" {
			log.Info("empty textinput")
			statusCode = http.StatusBadRequest
			// result = "Type or paste your art here..."
			HTTPResponseHeader = "400: Bad Request"
		} else {

			arts := helpers.ParseArt(strings.Split(input, "\n")) // TODO: add error check and restructure the handler to fail early.

			for _, a := range arts {
				sb.WriteString(artProc.Encode(a))
				// sb.WriteByte('\n')
			}

			result = sb.String()

			if strings.Contains(result, "Error") {
				statusCode = http.StatusBadRequest
				HTTPResponseHeader = "400: Bad Request"
				log.Info("input contains malformed line")
			} else {
				statusCode = http.StatusAccepted
				HTTPResponseHeader = "202: Accepted"
				log.Info("input is valid")
			}
		}

		data := map[string]string{
			"input":        input,
			"result":       result,
			"HTTPResponse": HTTPResponseHeader,
		}
		var buf bytes.Buffer

		if err := indexTmpl.Execute(&buf, data); err != nil {
			log.Error("failed to execute template", "error", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)

			return
		}

		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(statusCode)
		buf.WriteTo(w)
	})
}
