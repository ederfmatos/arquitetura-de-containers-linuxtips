package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	http.HandleFunc("GET /health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		fmt.Fprintf(w, `{"ok": true}`)
	})

	http.HandleFunc("GET /user", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Chuck Norris"
		}
		slog.Info("Request received", "path", r.URL.Path, "name", name)
		fmt.Fprintf(w, `{"name": "%s"}`, name)
	})

	slog.Info("Server running")
	http.ListenAndServe(":8080", nil)
}
