package types

// JamRequestParams struct, is the struct
// on what the params from creating  a new jam
// are modeled into
type JamRequestParams struct {
	UserID   string  `json:"user_id"`
	Name     string  `json:"name"`
	Location string  `json:"location"`
	Lat      float64 `json:"lat,omitempty"`
	Lng      float64 `json:"lng,omitempty"`
	Notes    string  `json:"notes,omitempty"`
}

// UpdateJamRequestParams struct, is the struct
// on what the params from updating a jam
// are modeled into
type UpdateJamRequestParams struct {
	ID       string `json:"id"`
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
	Notes    string `json:"notes,omitempty"`
}

//JoinJamRequestParams struct, is the struct
// on what the params from joining a jam
// are modeled into
type JoinJamRequestParams struct {
	Pin    string `json:"pin"`
	UserID string `json:"user_id"`
}
type UploadJamParams struct {
	UserID      string
	FileName    string
	TempFileURL string
	JamID       string
	StartTime   string
	EndTime     string
}

//UserRequestParams struct, is the struct
// on what the params from registering a user
// are modeled into
type UserRequestParams struct {
	FBID        string `json:"facebook_id"`
	AccessToken string `json:"access_token"`
}
