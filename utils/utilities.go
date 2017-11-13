package utils

import (
	"crypto/rand"
	"dsound/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// ParseJam func, parses the incoming
// params from the create a new jam request
func ParseJam(r *http.Request) (types.JamRequestParams, error) {
	var p types.JamRequestParams
	userId := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		p.UserID = userId
		return p, nil
	}
	fmt.Println(err)
	return p, err
}

// ParseParams Parses the request parameters
// by passing the type you want to parse into
func ParseParams(r *http.Request, to interface{}) (interface{}, error) {
	p := to
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		return p, nil
	}
	return p, err
}

// ParseUser func, parses the incoming
// params from create a new user request
func ParseUser(r *http.Request) (types.UserRequestParams, error) {
	var p types.UserRequestParams
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		return p, nil
	}
	return types.UserRequestParams{}, err
}

func ParseJoinJam(r *http.Request) (types.JoinJamRequestParams, error) {
	var p types.JoinJamRequestParams
	userId := r.Header.Get("user_id")
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		p.UserID = userId
		return p, nil
	}
	return types.JoinJamRequestParams{}, err
}

func ParseUpload(r *http.Request) (types.UploadJamParams, error) {

	infile, _, err := r.FormFile("filename")

	if err != nil {

		return types.UploadJamParams{}, err
	}
	userID := r.Header.Get("user_id")
	jamID := r.FormValue("jam_id")
	startTime := r.FormValue("start_time")
	endTime := r.FormValue("end_time")
	fileName := r.FormValue("filename")

	p := types.UploadJamParams{
		UserID:      userID,
		FileName:    fileName,
		JamID:       jamID,
		StartTime:   startTime,
		EndTime:     endTime,
		TempFileURL: ".uploads/" + userID,
	}
	outfile, err := os.Create(".uploads/" + userID)
	if err != nil {

		return types.UploadJamParams{}, err
	}

	_, err = io.Copy(outfile, infile)
	if err != nil {

		return types.UploadJamParams{}, err
	}

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
