package routes

import (
	"github.com/gin-gonic/gin"
)

//MainRouter is the main Gin router
// instance
var MainRouter = gin.Default()

const (
	//APIV is the current version string for
	// our api
	APIV = "/api/v2.0/"
)

func init() {
	// TODO:
}
func index(c *gin.Context) {
	c.JSON(200, gin.H{"index route": "main index route not handle"})
}

// StartServer will start the server on
// port 8080 and add all the sub routes
func StartServer() {
	addToMainRouter()
	addUserRoutes()
	MainRouter.GET("/", index)
	MainRouter.Run(":8080")
}

// func setContentTypeJSON(next httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		w.Header().Set("Content-Type", "application/json")
// 		next(w, r, p)
// 	}
// }

// // Auth func, authenticates the user
// // if user is not authenticated then
// // it will retoute to bad request
// func Auth(next httprouter.Handle) httprouter.Handle {
// 	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
// 		usrID := r.Header.Get("user_id")
// 		if usrID == "" {
// 			w.WriteHeader(http.StatusUnauthorized)
// 			json.NewEncoder(w).Encode(types.ResponseMessage{M: "Not authorized"})
// 			return
// 		}
// 		next(w, r, p)
// 	}
// }
