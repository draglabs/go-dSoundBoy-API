package routes

import (
	"github.com/draglabs/go-dSoundBoy-API/controllers"
	"github.com/draglabs/go-dSoundBoy-API/types"
	"github.com/draglabs/go-dSoundBoy-API/utils"

	"fmt"

	"github.com/gin-gonic/gin"
)

//JamRouter struct, is the jam router
var jamRouter = MainRouter.Group(APIV + "jam/")

const (
	recordingsR = "recording/:id"
	jamNewR     = "new"
	joinR       = "join"
	upload      = "upload"
	details     = "details/:id"
	updateJamR  = "update"
)

// addToMainROuter func, will add all the jam routes
// to he main router
func addToMainRouter() {
	jamRouter.GET(recordingsR, recordings)
	jamRouter.GET(details, jamDetails)
	jamRouter.POST(jamNewR, newJam)
	jamRouter.POST(joinR, join)
	jamRouter.POST(upload, uploadAudioFile)
	jamRouter.POST(updateJamR, updateJam)
}

// new func, will give us a new jam regarless of the user having an
// active jam, if the user has an active jam it will be replaced by this one
func newJam(c *gin.Context) {
	pm, err := utils.ParseJam(c)
	if err != nil {
		c.JSON(400, types.ResponseMessage{M: "One or more params are missing"})
		return
	}
	jam, err := controllers.Jam.Create(pm)
	if err == nil {
		c.JSON(200, jam)
		return
	}
	c.JSON(500, types.ResponseMessage{M: "Unable to create Jam"})
}

// upload func, takes care of the uploading, and currently uploads the file to
// s3 bucket.
func uploadAudioFile(c *gin.Context) {
	para, err := utils.ParseUpload(c)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something went wrong"})
		return
	}
	err = controllers.Jam.Upload(para)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something went wrong"})
		return
	}
	c.JSON(200, types.ResponseMessage{M: "uploaded successfully"})
}

// join func, join a user into a jam.
func join(c *gin.Context) {
	para, err := utils.ParseJoinJam(c)
	if err != nil {
		fmt.Println("error parsing join jam " + err.Error())
		c.JSON(500, types.ResponseMessage{M: "One or more params are missing"})
		return
	}

	if jam, err := controllers.Jam.Join(para); err == nil {
		c.JSON(200, jam)
		return
	}
}

func jamDetails(c *gin.Context) {
	jam, err := controllers.Jam.Details(c.Param("id"))
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "Something when wrong, Error: " + err.Error()})
		return
	}
	c.JSON(200, jam)

}

// recordings func, will fetch all the recordings for a given jam id.
func recordings(c *gin.Context) {
	id := c.Param("id")
	recordings, err := controllers.Recordings(id)
	if err != nil {
		c.JSON(400, types.ResponseMessage{M: "No recordings for this jam " + id})
		return
	}
	c.JSON(200, recordings)
}
func updateJam(c *gin.Context) {
	para, err := utils.ParseUpdate(c)
	jam, err := controllers.Jam.Update(para)
	if err != nil {
		c.JSON(500, types.ResponseMessage{M: "something went wrong" + err.Error()})
		return
	}
	if err == nil {
		c.JSON(200, jam)
	}
}
