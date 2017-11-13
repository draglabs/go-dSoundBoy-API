package routes

import (
	"dsound/types"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	APIV = "/api/v1.0/"
)

// Router is the main router
// on which sub routes are registered
var Router = httprouter.New()

func AddAllSubRoutes() {
	jam := NewJamRouter()
	jam.addToMainRouter(Router)
	Router.GET("/", index)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode(types.ResponseMesssage{"homepage route"})
}
