package db

import (
	"gopkg.in/mgo.v2"
)

func EnsureIndices() {
	db := NewDB()
	defer db.Close()

	db.JamCollection().EnsureIndex(mgo.Index{Key: []string{"pin"}, Unique: true, Name: "pin_index"})
	db.RecordingsCollection().EnsureIndex(mgo.Index{Key: []string{"jam_id", "user_id"}})
	db.UserCollection().EnsureIndex(mgo.Index{Key: []string{"fb_id"}})
}
