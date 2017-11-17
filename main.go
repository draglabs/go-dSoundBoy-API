package main

import (
	"dsound/db"
	"dsound/routes"
	"log"
	"net/http"
)

func main() {
	db.EnsureIndices()
	routes.AddAllSubRoutes()
	log.Fatal(http.ListenAndServe(":8080", routes.Router))
}
