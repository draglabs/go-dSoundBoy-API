package utils

import (
	"dsound/types"
	"encoding/json"
	"net/http"
)

// ParseJam func, parses the incoming
// params from the create a new jam request
func ParseJam(r *http.Request) (types.JamRequestParams, error) {
	var p types.JamRequestParams
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		return p, nil
	}
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
	err := json.NewDecoder(r.Body).Decode(&p)
	if err == nil {
		return p, nil
	}
	return types.JoinJamRequestParams{}, err
}
