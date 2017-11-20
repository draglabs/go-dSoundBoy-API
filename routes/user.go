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
	register      = userBaseRoute + "/register"
	activity      = userBaseRoute + "/activity"
	activeJam     = userBaseRoute + "/jam/active"
	update        = userBaseRoute + "/update"
)

type UserRouter struct {
}

func NewUserRouter() UserRouter {
	return UserRouter{}
}
func (ur *UserRouter) AddUserRoutes(r *httprouter.Router) {
	r.POST(register, setContentTypeJSON(ur.register))
	r.GET(activeJam, setContentTypeJSON(ur.activeJam))
	r.GET(activity, setContentTypeJSON(ur.activity))
	r.PUT(update, setContentTypeJSON(ur.update))
}
func (ur *UserRouter) register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pa, err := utils.ParseCreateUser(r)
	if err == nil {
		usr, _ := controllers.User.Register(pa)
		json.NewEncoder(w).Encode(usr)
		return
	}

}

func (ur *UserRouter) activeJam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pa := utils.ParseUserID(r)
	jam, err := controllers.User.ActiveJam(pa)
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "Cant Find Active Jam"})
		return
	}
	json.NewEncoder(w).Encode(jam)
}
func (ur *UserRouter) activity(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	jams, err := controllers.User.Activity(utils.ParseUserID(r))
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "Unable to find user activity"})
		return
	}
	json.NewEncoder(w).Encode(jams)
}

func (ur *UserRouter) update(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	user, err := controllers.User.Update(utils.ParseUserID(r))
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "Unable to update user error: " + err.Error()})
		return
	}
	json.NewEncoder(w).Encode(user)
}
