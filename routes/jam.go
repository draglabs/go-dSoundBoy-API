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
	jamR        = APIV + "jam"
	jamByID     = jamR + "/:id"
	recordingsR = jamR + "/recording/:id"
	jamNewR     = jamR + "/new"
	joinR       = jamR + "/join"
	upload      = jamR + "/upload"
	details     = jamR + "/details/:id"
)

//NewJamRouter func, gives us a new JamRouter
func NewJamRouter() JamRouter {
	return JamRouter{}
}

// addToMainROuter func, will add all the jam routes
// to he main router
func (j *JamRouter) addToMainRouter(r *httprouter.Router) {
	r.GET(jamByID, setContentTypeJSON(j.jam))
	r.GET(details, setContentTypeJSON(j.details))
	r.POST(jamNewR, setContentTypeJSON(j.new))
	r.POST(joinR, setContentTypeJSON(j.join))
	r.POST(upload, j.upload)
}

// jam func, fetches a jam by id
func (j *JamRouter) jam(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	jm, err := controllers.Jam.FindByID(id)
	if err == nil {
		json.NewEncoder(w).Encode(jm)
		return
	}

}

// new func, will give us a new jam regarless of the user having an
// active jam, if the user has an active jam it will be replaced by this one
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

// upload func, takes care of the uplaoding, and currently uploads the file to
// s3 bucket.
func (j *JamRouter) upload(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	para, err := utils.ParseUpload(r)
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "Something when wrong"})
		return
	}
	err = controllers.Jam.Upload(para)
	if err == nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "uploaded succesfuly"})
	}
}

// join func, join a user into a jam.
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

func (j *JamRouter) details(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	jam, err := controllers.Jam.Details(p.ByName("id"))
	if err == nil {
		json.NewEncoder(w).Encode(jam)
		return
	}
	json.NewEncoder(w).Encode(types.ResponseMessage{M: "Something when wrong, Error: " + err.Error()})

}

// recordings func, will fetch all the recordings for a given jam id.
func (j *JamRouter) recordings(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")
	recordings, err := controllers.Recordings(id)
	if err != nil {
		json.NewEncoder(w).Encode(types.ResponseMessage{M: "No recordings for this jam " + id})
		return
	}
	json.NewEncoder(w).Encode(recordings)
}
