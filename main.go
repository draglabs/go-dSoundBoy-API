package main

import (
	"dsound/routes"
	"net/http"
)

func main() {
	routes.AddAllSubRoutes()
	http.ListenAndServe(":80", routes.Router)
}
