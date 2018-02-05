package types

// JamResponse struct is the a light
// struct of the jam object
type JamResponse struct {
	ID            string    `json:"id"`
	Name          string    `json:"name"`
	StartTime     string    `json:"start_time"`
	EndtTime      string    `json:"end_time"`
	Location      string    `json:"location"`
	Notes         string    `json:"notes"`
	Collaborators int       `json:"collaborators"`
	Link          string    `json:"link"`
	Pin           string    `json:"pin"`
	Coordinates   []float64 `json:"coordinates"`
}

type UserResponse struct {
}
type ResponseMessage struct {
	M string `json:"message"`
}
