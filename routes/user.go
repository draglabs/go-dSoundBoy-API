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
