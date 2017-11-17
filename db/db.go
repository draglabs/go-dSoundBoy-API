package db

import mgo "gopkg.in/mgo.v2"
import "fmt"

func ensureIndices() {

}

const (
	DB         = "dsound"
	jam        = "jams"
	user       = "users"
	recordings = "recordings"
)

func NewDB() DataStore {

	info := mgo.DialInfo{
		Addrs:    []string{"54.183.100.139:27017"},
		Database: DB,
		Username: "soundBoy",
		Password: "soundBoy",
	}
	s, err := mgo.DialWithInfo(&info)
	fmt.Println(err)

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
	c := d.session.DB(DB).C(jam)
	return c
}

func (d DataStore) UserCollection() *mgo.Collection {
	c := d.session.DB(DB).C(user)
	return c
}

func (d DataStore) RecordingsCollection() *mgo.Collection {
	c := d.session.DB(DB).C(recordings)
	return c
}
