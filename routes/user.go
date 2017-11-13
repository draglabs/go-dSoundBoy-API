package routes

import (
	"dsound/controllers"
	"dsound/types"
	"dsound/utils"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

const (
	userBaseRoute = APIV + "user"
	register      = userBaseRoute + "register"
	activity      = userBaseRoute + "/activity"
	activeJam     = userBaseRoute + "/jam/active"
)

type UserRouter struct {
}

func NewUserRouter() UserRouter {
	return UserRouter{}
}
func (ur *UserRouter) AddUserRoutes(r *httprouter.Router) {
	r.POST(register, setContentTypeJSON(ur.register))
	r.GET(activeJam, setContentTypeJSON(ur.activeJam))
}
func (ur *UserRouter) register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	para, err := utils.ParseUser(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "One or more params are missing"})
	}
	if usr, err := controllers.User.Register(para); err == nil {
		json.NewEncoder(w).Encode(usr)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(types.ResponseMessage{M: "Unable to register user"})
}

func (ur *UserRouter) activeJam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
