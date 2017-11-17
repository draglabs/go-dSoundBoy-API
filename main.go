package main

import (
	"dsound/routes"
	"net/http"
	"os"
)

func main() {
	routes.AddAllSubRoutes()
	http.ListenAndServe(os.Getenv("PORT"), routes.Router)
}
