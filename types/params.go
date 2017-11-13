package types

type JamRequestParams struct {
	Name     string  `json:"name"`
	Location string  `json"location"`
	Lat      float64 `json:"latitude"`
	Lng      float64 `json:"longitude"`
	Notes    string  `json:notes,omitemty"`
}

type UpdateJamRequestParams struct {
	ID       string `json:"id`
	Name     string `json:"name,omitempty"`
	Location string `json"location,omitempty"`
	Notes    string `json:notes,omitemty"`
}

type JoinJamRequestParams struct {
	ID       string `json:"id`
	Name     string `json:"name,omitempty"`
	Location string `json"location,omitempty"`
	Notes    string `json:notes,omitemty"`
}

type UserRequestParams struct {
	FB string `json:"fb`
}
