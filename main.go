package main

import (
	"embed"
	"io/fs"
	"log/slog"
	"net/http"
	"template/pages"
	"time"

	"github.com/go-chi/chi/v5"
)

var (
	//go:embed static/*
	staticFS embed.FS
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home().Render(r.Context(), w)
	})

	staticSubFS, err := fs.Sub(staticFS, "static")
	if err != nil {
		panic(err)
	}
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServerFS(staticSubFS)))

	s := http.Server{
		Addr:         ":3000",
		Handler:      r,
		ReadTimeout:  time.Second * 10,
		WriteTimeout: time.Second * 30,
		IdleTimeout:  time.Second * 5,
	}

	slog.Info("Listening on http://localhost:3000")
	s.ListenAndServe()
}
