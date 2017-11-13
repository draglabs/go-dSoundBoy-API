package controllers

import (
	"dsound/db"
	"dsound/models"
	"dsound/types"

	"gopkg.in/mgo.v2/bson"
)

type jam struct {
}

func newJam() jam {
	return jam{}
}

var Jam = newJam()

func (j jam) Create(p types.JamRequestParams) (models.Jam, error) {
	id := bson.NewObjectId().String()

	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	jam := models.Jam{
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
	err := c.FindId(id).One(&jm)
	if err == nil {
		return jm, nil
	}
	return jm, err
}
