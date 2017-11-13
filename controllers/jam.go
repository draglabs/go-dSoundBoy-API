package controllers

import (
	"crypto/rand"
	"dsound/db"
	"dsound/models"
	"dsound/types"
	"io"

	"gopkg.in/mgo.v2/bson"
)

type jam struct {
}

func newJam() jam {
	return jam{}
}

var Jam = newJam()

func (j jam) Create(p types.JamRequestParams) (models.Jam, error) {
	id := bson.NewObjectId()

	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	jam := models.Jam{
		Pin:         encodeToString(4),
		ID:          id,
		Name:        p.Name,
		Location:    p.Location,
		Coordinates: []float64{p.Lat, p.Lng},
	}
	err := c.Insert(jam)
	if err == nil {
		return jam, nil
	}
	return jam, err
}

func (j jam) upload() {

}
func (j jam) Join() {

}
func (j jam) Update(p types.UpdateJamRequestParams) error {
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	err := c.Update(p.ID, p)
	if err != nil {
		return err
	}
	return nil
}

func (j jam) FindId(id string) (models.Jam, error) {
	var jm models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	err := c.FindId(bson.ObjectIdHex(id)).One(&jm)
	if err == nil {
		return jm, nil
	}
	return jm, err
}
func encodeToString(max int) string {
	var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
	b := make([]byte, max)
	n, err := io.ReadAtLeast(rand.Reader, b, max)
	if n != max {
		panic(err)
	}
	for i := 0; i < len(b); i++ {
		b[i] = table[int(b[i])%len(table)]
	}
	return string(b)
}
