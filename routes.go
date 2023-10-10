package brimstoneesan

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (b *Brimstoneesan) routes() http.Handler {
	mux := chi.NewRouter()
	mux.Use(middleware.RequestID)
	mux.Use(middleware.RealIP)
	mux.Use(middleware.Logger)
	if b.Debug {
	}
	mux.Use(middleware.Recoverer)
	mux.Use(b.SessionLoad)

	return mux
}
