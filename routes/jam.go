package routes

import (
	"dsound/controllers"
	"dsound/types"
	"dsound/utils"
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//JamRouter struct, is the jam router
type JamRouter struct {
}

const (
	jamR    = APIV + "jam"
	jamNewR = jamR + "/new"
	joinR   = jamR + "/join"
	upload  = jamR + "/upload"
)

//NewJamRouter func, gives us a new JamRouter
func NewJamRouter() JamRouter {
	return JamRouter{}
}

func (j *JamRouter) addToMainRouter(r *httprouter.Router) {
	r.GET(jamR, setContentTypeJSON(j.jam))
	r.POST(jamNewR, setContentTypeJSON(j.new))
	r.POST(joinR, setContentTypeJSON(j.join))
	r.POST(upload, j.upload)
}

func (j *JamRouter) jam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	r.ParseForm()
	id := r.FormValue("id")
	jm, err := controllers.Jam.FindById(id)
	if err == nil {
		json.NewEncoder(w).Encode(jm)
		return
	}

}

func (j *JamRouter) new(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	pm, err := utils.ParseJam(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "One or more params are missing"})
		return
	}
	jam, err := controllers.Jam.Create(pm)
	if err == nil {
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(jam)
		return
	}
	w.WriteHeader(http.StatusInternalServerError)
	json.NewEncoder(w).Encode(types.ResponseMessage{M: "Unable to create Jam"})
}

func (j *JamRouter) upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	para, err := utils.ParseUpload(r)
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "Something when wrong"})
		return
	}
	err = controllers.Jam.Upload(para)
	if err == nil {

	}
}

func (j *JamRouter) join(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	para, err := utils.ParseJoinJam(r)
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "One or more params are missing"})
		return
	}
	if jam, err := controllers.Jam.Join(para); err == nil {
		json.NewEncoder(w).Encode(jam)
	}
}

func (j *JamRouter) recordings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

}
