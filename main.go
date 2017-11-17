package main

import (
	"dsound/routes"
	"log"
	"net/http"
)

func main() {
	routes.AddAllSubRoutes()
	log.Fatal(http.ListenAndServe(":80", routes.Router))
}
