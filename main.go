package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/patrickmn/go-cache"

	"github.com/Jeswyrne/chlnge/pkg/user"
)

func main() {
	router := chi.NewRouter()

	// Cache
	cache := cache.New(cache.DefaultExpiration, cache.DefaultExpiration)

	user := user.NewUser(cache)

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Get("/users/info", user.Handler)

	http.ListenAndServe(":3000", router)
}
