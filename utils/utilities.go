package utils

import (
	"crypto/rand"
	"dsound/types"
	"fmt"
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

// ParseJam func, parses the incoming
// params from the create a new jam request
func ParseJam(c *gin.Context) (types.JamRequestParams, error) {
	var p types.JamRequestParams
	userID := c.GetHeader("user_id")
	err := c.Bind(&p)
	if err == nil {
		p.UserID = userID
		return p, nil
	}
	return p, err
}

// ParseUpdate func
func ParseUpdate(c *gin.Context) (types.UpdateJamRequestParams, error) {
	var p types.UpdateJamRequestParams
	err := c.Bind(&p)
	return p, err
}

// ParseCreateUser func, parses the incoming
// params from create a new user request
func ParseCreateUser(c *gin.Context) (types.CreateUserParams, error) {
	var p types.CreateUserParams

	err := c.Bind(&p)
	return p, err
}

// ParseUserID func
func ParseUserID(c *gin.Context) string {
	id := c.GetHeader("user_id")
	return id
}

// ParseJoinJam func
func ParseJoinJam(c *gin.Context) (types.JoinJamRequestParams, error) {
	var p types.JoinJamRequestParams
	//userId := r.Header.Get("user_id")
	err := c.Bind(&p)
	return p, err
}

// ParseUpload func
func ParseUpload(c *gin.Context) (types.UploadJamParams, error) {

	// infile, _, err := r.FormFile("audioFile")
	infile, err := c.FormFile("audioFile")
	if err != nil {
		fmt.Println("in file error", err)
		return types.UploadJamParams{}, err
	}
	// userID := r.FormValue("user_id") // replaced for now., should come on the header.
	// jamID := r.FormValue("id")
	// startTime := r.FormValue("start_time")
	// endTime := r.FormValue("end_time")

	userID := c.PostForm("user_id")
	jamID := c.PostForm("id")
	startTime := c.PostForm("start_time")
	endTime := c.PostForm("end_time")

	p := types.UploadJamParams{
		UserID:      userID,
		JamID:       jamID,
		StartTime:   startTime,
		EndTime:     endTime,
		TempFileURL: ".uploads/" + jamID,
	}
	outfile, err := os.Create(".uploads/" + jamID)
	if err != nil {
		fmt.Println("outfile error", err)
		return types.UploadJamParams{}, err
	}
	fl, _ := infile.Open()
	_, err = io.Copy(outfile, fl)
	if err != nil {
		fmt.Println(err)
		return types.UploadJamParams{}, err
	}
	return p, err
}

// GeneratePin func
func GeneratePin(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
