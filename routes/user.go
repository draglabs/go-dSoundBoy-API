package routes

import (
	"dsound/controllers"
	"dsound/types"
	"dsound/utils"

	"github.com/gin-gonic/gin"
)

const (
	userBaseRoute = APIV + "user/"
	register      = userBaseRoute + "register"
	activity      = userBaseRoute + "activity"
	activeJam     = userBaseRoute + "jam/active"
	update        = userBaseRoute + "update"
)

//var userRouter = MainRouter.Group("/user/")
var ur = MainRouter.Group(userBaseRoute)

func addUserRoutes() {
	ur.POST(register, registerUser)
	ur.GET(activeJam, userActiveJam)
	ur.GET(activity, userActivity)
	ur.PUT(update, updateUser)
}

func registerUser(c *gin.Context) {
	pa, err := utils.ParseCreateUser(c)
	if err == nil {
		usr, _ := controllers.User.Register(pa)
		c.JSON(200, usr)
	}

}

func userActiveJam(c *gin.Context) {
	pa := utils.ParseUserID(c)
	jam, err := controllers.User.ActiveJam(pa)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Cant Find Active Jam"})
	}
	c.JSON(200, jam)
}

func userActivity(c *gin.Context) {

	jams, err := controllers.User.Activity(utils.ParseUserID(c))
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Unable to find user activity"})
	}
	c.JSON(200, jams)
}

func updateUser(c *gin.Context) {
	user, err := controllers.User.Update(utils.ParseUserID(c))
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Unable to update user error: " + err.Error()})
	}
	c.JSON(200, user)
}
