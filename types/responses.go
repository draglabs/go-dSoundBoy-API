package types

type JamResponse struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	StartTime string `json:"start_time"`
	EndtTime  string `json:"end_time"`
	Location  string `json:"location"`
	Notes     string `json:"notes"`
}

type UserResponse struct {
}
type ResponseMessage struct {
	M string `json:"message"`
}
