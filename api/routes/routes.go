package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patrickmn/go-cache"

	"github.com/Jeswyrne/chlnge/api/controller"
	"github.com/Jeswyrne/chlnge/api/middlewares"
)

func Routes() *chi.Mux {
	r := chi.NewRouter()

	// Cache
	cache := cache.New(cache.DefaultExpiration, cache.DefaultExpiration)

	// User Controller
	userController := controller.NewUser(cache)

	// Middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.RedirectSlashes)

	// Custom Middleware
	r.Use(middlewares.SetMiddlewareHeaders)

	r.Get("/users/info", userController.Handler)

	return r
}
