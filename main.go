package main

import (
	"log/slog"
	"net/http"
	"template/pages"
	"time"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		pages.Home().Render(r.Context(), w)
	})

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
