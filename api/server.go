package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jeswyrne/chlnge/api/routes"
)

func Run() {
	port := ":3000"

	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(port, routes.Routes()))
}
