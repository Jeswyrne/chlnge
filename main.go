package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	router := chi.NewRouter()

	// Middleware
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)
	router.Use(middleware.RedirectSlashes)

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {

		users := r.URL.Query().Get("users")
		userList := strings.Split(users, ",")

		fmt.Print(userList)
		w.Write([]byte("pong"))
	})

	http.ListenAndServe(":3333", router)
}
