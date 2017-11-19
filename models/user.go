package models

//User struct, is the struct
// that an user is modeled into
type User struct {
	ID         string `json:"id" bson:"_id"`
	FirstName  string `json:"name" bson:"first_name"`
	LastName   string `json:"last_name" bson:"last_name"`
	FBEmail    string `json:"fb_email" bson:"fb_email"`
	FBID       string `json:"fb_id" bson:"fb_id"`
	Email      string `json:"email,omitempty" bson:"email,omitempty"`
	CurrentJam *Jam   `json:"current_jam" bson:"current_jam"`
}
