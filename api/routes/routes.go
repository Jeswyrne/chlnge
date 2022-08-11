package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patrickmn/go-cache"

	"github.com/Jeswyrne/chlnge/api/controller"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	// Cache
	cache := cache.New(cache.DefaultExpiration, cache.DefaultExpiration)

	userController := controller.NewUser(cache)

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	r.Get("/users/info", userController.Handler)

	return r
}
