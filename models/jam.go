package models

// Jam struct, models our Jam collection
type Jam struct {
	ID     string `json:"id" bson:"_id"`
	Pin    string `json:"pin"    bson:"pin"`
	Status bool   `json:"status" bson:"status"`
	Name   string `json:"name"   bson:"name"`

	Coordinates   []float64    `json:"coordinates"     bson:"coordinates"`
	Collaborators []Creator    `json:"collaborators"   bson:"collaborators"`
	Recordings    []Recordings `json:"recordings"      bson:"recordings"`
	Location      string       `json:"location"        bson:"location"`
	Creator       Creator      `json:"creator"         bson:"creator"`
	StartTime     string       `json:"start_time"      bson:"start_time"`
	EndTime       string       `json:"end_time"        bson:"end_time"`
	StatusID      string       `json:"status_id"       bson:"status_id"`
	Notes         string       `json:"notes"           bson:"notes"`
}

type Recordings struct {
	User      Creator `json:"user" bson:"user"`
	FileName  string  `json:"file_name" bson:"file_name"`
	JamID     string  `json:"jam_id" bson:"jam_id"`
	StartTime string  `json:"start_time" bson:"start_time"`
	EndTime   string  `json:"end_time" bson:"end_time"`
	Notes     string  `json:"notes" bson:"notes"`
	S3url     string  `json:"s3url" bson:"s3url"`
}

//Creator struct
type Creator struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}
