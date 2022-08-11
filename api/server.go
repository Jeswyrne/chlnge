package api

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Jeswyrne/chlnge/api/routes"
)

func Run() {
	port := ":3000"

	fmt.Println("Listening to port 3000")
	log.Fatal(http.ListenAndServe(port, routes.Routes()))
}
