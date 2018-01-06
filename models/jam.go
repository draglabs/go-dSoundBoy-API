package models

// Jam struct, is the struct on which
// a Jam is modeled into
type Jam struct {
	ID            string       `json:"id"             bson:"_id"`
	Pin           string       `json:"pin"             bson:"pin"`
	IsCurrent     bool         `json:"is_current"      bson:"is_current"`
	Name          string       `json:"name"            bson:"name"`
	UserID        string       `json:"user_id"         bson:"user_id"`
	Coordinates   []float64    `json:"coordinates"     bson:"coordinates"`
	Collaborators []User       `json:"collaborators"   bson:"collaborators"`
	Recordings    []Recordings `json:"recordings"      bson:"recordings"`
	Location      string       `json:"location"        bson:"location"`
	StartTime     string       `json:"start_time"      bson:"start_time"`
	EndTime       string       `json:"end_time"        bson:"end_time"`
	Notes         string       `json:"notes"           bson:"notes"`
	Link          string       `json:"link" bson:"link"`
}

// Recordings struct, is the struct on which
// the recordings for a jam are modeled into
type Recordings struct {
	ID        string `json:"id" bson:"_id"`
	UserID    string `json:"user_id" bson:"user_id"`
	FileName  string `json:"file_name"   bson:"file_name"`
	JamID     string `json:"jam_id"      bson:"jam_id"`
	StartTime string `json:"start_time"  bson:"start_time"`
	EndTime   string `json:"end_time"    bson:"end_time"`
	Notes     string `json:"notes"       bson:"notes"`
	S3url     string `json:"s3url"       bson:"s3url"`
}
