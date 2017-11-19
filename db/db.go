package db

import mgo "gopkg.in/mgo.v2"

const (
	db         = "dsound"
	jam        = "jams"
	user       = "users"
	recordings = "recordings"
)

// NewDB func, give us a new mgo session
func NewDB() DataStore {

	info := mgo.DialInfo{
		Addrs:    []string{"54.183.100.139:27017"},
		Database: db,
		Username: "soundBoy",
		Password: "soundBoy",
	}
	s, _ := mgo.DialWithInfo(&info)

	s.SetMode(mgo.Monotonic, true)
	return DataStore{
		session: s.Copy(),
	}
}
func (d DataStore) Close() {
	d.session.Close()
}

type DataStore struct {
	session *mgo.Session
}

func (d DataStore) JamCollection() *mgo.Collection {
	c := d.session.DB(db).C(jam)
	return c
}

func (d DataStore) UserCollection() *mgo.Collection {
	c := d.session.DB(db).C(user)
	return c
}

func (d DataStore) RecordingsCollection() *mgo.Collection {
	c := d.session.DB(db).C(recordings)
	return c
}
