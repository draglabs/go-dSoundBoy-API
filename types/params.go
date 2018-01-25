package types

// JamRequestParams struct, is the struct
// on what the params from creating  a new jam
// are modeled into
type JamRequestParams struct {
	UserID   string  `form:"user_id"`
	Name     string  `form:"name"`
	Location string  `form:"location"`
	Lat      float64 `form:"lat,omitempty"`
	Lng      float64 `form:"lng,omitempty"`
	Notes    string  `form:"notes,omitempty"`
}

// UpdateJamRequestParams struct, is the struct
// on what the params from updating a jam
// are modeled into
type UpdateJamRequestParams struct {
	ID       string `form:"id" bson:"-"`
	Name     string `form:"name,omitempty" bson:"name,omitempty"`
	Location string `form:"location,omitempty" bson:"location,omitempty"`
	Notes    string `form:"notes,omitempty" bson:"notes,omitempty"`
}

//JoinJamRequestParams struct, is the struct
// on what the params from joining a jam
// are modeled into
type JoinJamRequestParams struct {
	Pin    string `form:"pin"`
	UserID string `form:"user_id"`
}
type UploadJamParams struct {
	UserID      string
	FileName    string
	TempFileURL string
	JamID       string
	StartTime   string
	EndTime     string
}

//CreateUserParams struct, is the struct
// on what the params from registering a user
// are modeled into
type CreateUserParams struct {
	FBID        string `form:"facebook_id"`
	AccessToken string `form:"access_token"`
}
type UpdateUserParams struct {
	Email string `form:"email"`
}
