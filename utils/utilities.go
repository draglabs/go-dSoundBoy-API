package utils

import (
	"crypto/rand"
	"dsound/types"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
)

// ParseJam func, parses the incoming
// params from the create a new jam request
func ParseJam(r *http.Request) (types.JamRequestParams, error) {
	var p types.JamRequestParams
	userID := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&p)
	defer r.Body.Close()
	if err == nil {
		p.UserID = userID
		return p, nil
	}
	fmt.Println(err)
	return p, err
}
func ParseUpdate(r *http.Request) (types.UpdateJamRequestParams, error) {
	var p types.UpdateJamRequestParams
	userID := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&p)
	defer r.Body.Close()
	if err == nil {
		p.ID = userID
		return p, nil
	}
	fmt.Println(err)
	return p, err
}

// ParseCreateUser func, parses the incoming
// params from create a new user request
func ParseCreateUser(r *http.Request) (types.CreateUserParams, error) {
	var p types.CreateUserParams

	b, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	err = json.Unmarshal(b, &p)

	return p, err
}

func ParseUserID(r *http.Request) string {
	id := r.Header.Get("user_id")
	return id

}
func ParseJoinJam(r *http.Request) (types.JoinJamRequestParams, error) {
	var p types.JoinJamRequestParams
	userId := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&p)
	defer r.Body.Close()
	if err == nil {
		p.UserID = userId
		return p, nil
	}
	return types.JoinJamRequestParams{}, err
}

func ParseUpload(r *http.Request) (types.UploadJamParams, error) {

	infile, _, err := r.FormFile("audioFile")

	if err != nil {
		fmt.Println(err)
		return types.UploadJamParams{}, err
	}
	userID := r.FormValue("user_id") // replaced for now., whould come on the header.
	jamID := r.FormValue("id")
	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")
	fileName := r.FormValue("name")

	p := types.UploadJamParams{
		UserID:      userID,
		FileName:    fileName,
		JamID:       jamID,
		StartTime:   startTime,
		EndTime:     endTime,
		TempFileURL: ".uploads/" + jamID,
	}
	fmt.Println("userid", jamID)
	outfile, err := os.Create(".uploads/" + jamID)
	if err != nil {
		fmt.Println(err)
		return types.UploadJamParams{}, err
	}

	_, err = io.Copy(outfile, infile)
	if err != nil {
		fmt.Println(err)
		return types.UploadJamParams{}, err
	}
	fmt.Println("params from upload", p)
	return p, err
}

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
