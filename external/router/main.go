package router

import (
	"context"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

type (
	Mux struct {
		*chi.Mux
	}
)

func New(ctx context.Context) (*Mux, error) {

	// New router instance
	router := &Mux{
		chi.NewRouter(),
	}

	// Add middlewares
	router.Use(
		// Enable CORS for all routes
		cors.AllowAll().Handler,
		// Set the content type to JSON for all responses
		render.SetContentType(render.ContentTypeJSON),
	)

	// Define routes for the router
	router.Route("/v1", func(r chi.Router) {
		// Mount the user handler on the "/v1/user" route
		r.Mount("/user", router.handlerUser())
	})

	return router, nil
}

func (mux *Mux) handlerUser() chi.Router {
	router := chi.NewRouter()

	// TODO: Implement user comtroller via manager
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return router
}
