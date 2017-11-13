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
	usr := NewUserRouter()
	jam.addToMainRouter(Router)
	usr.AddUserRoutes(Router)
	Router.GET("/", index)
}

func index(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	json.NewEncoder(w).Encode(types.ResponseMessage{M: "homepage route"})
}
func setContentTypeJSON(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		next(w, r, p)
	}
}

// Auth func, authenticates the user
// if user is not authenticated then
// it will retoute to bad request
func Auth(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		usrID := r.Header.Get("user_id")
		if usrID == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(types.ResponseMessage{M: "Not authorized"})
			return
		}
		next(w, r, p)
	}
}
