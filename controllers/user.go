package controllers

import (
	"dsound/db"
	"dsound/models"
	"dsound/types"
	"dsound/vendor"
	"fmt"
	"log"

	"gopkg.in/mgo.v2/bson"
)

type user struct {
}

var User = newUser()

func newUser() user {
	return user{}
}

func (u user) Register(p types.CreateUserParams) (models.User, error) {

	usr, err := u.FindByFB(p.FBID)
	if err == nil {
		return usr, nil
	}
	return createUser(p)
}

func createUser(p types.CreateUserParams) (models.User, error) {

	usr, err := vendor.FBFetchUser(p)

	if err == nil {
		db := db.NewDB()
		defer db.Close()
		c := db.UserCollection()
		err = c.Insert(usr)
		return usr, err
	}
	return usr, err

}

func (u user) FindByID(id string) (models.User, error) {
	var user models.User
	db := db.NewDB()
	defer db.Close()
	c := db.UserCollection()
	err := c.FindId(id).One(&user)
	if err == nil {
		return user, nil
	}
	return models.User{}, err
}

func (u user) FindByFB(fbID string) (models.User, error) {
	var user models.User
	db := db.NewDB()
	defer db.Close()
	c := db.UserCollection()
	err := c.Find(bson.M{"fb_id": fbID}).One(&user)
	if err == nil {
		return user, nil
	}
	return models.User{}, err
}

// UpdateCurrentJam func, will update the current jam
// for the user, if it cant update it, it will panic
// since this operation is key to the whole flow.
func (u user) UpdateCurrentJam(useID string, jam models.Jam) error {
	var user models.User
	var activeJam models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.UserCollection()
	errFindJam := db.JamCollection().Find(bson.M{"user_id": useID, "is_current": true}).One(&activeJam)
	if errFindJam == nil {
		activeJam.IsCurrent = false
		er := db.JamCollection().Update(bson.M{"_id": activeJam.ID}, bson.M{"$set": bson.M{"is_current": false}})
		fmt.Println(er)
	}
	fmt.Println(errFindJam)
	c.FindId(useID).One(&User)
	user.CurrentJam = &jam
	err := c.Update(bson.M{"_id": useID}, bson.M{"$set": bson.M{"current_jam": jam}})
	if err != nil {
		log.Fatal(err)
	}

	return err
}

func (u user) Activity(useID string) ([]types.JamResponse, error) {
	var response []types.JamResponse
	var jams []models.Jam
	db := db.NewDB()
	defer db.Close()
	jc := db.JamCollection()
	err := jc.Find(bson.M{"user_id": useID}).All(&jams)
	if err != nil {
		return response, err
	}
	for _, jm := range jams {
		resp := types.JamResponse{
			ID:        jm.ID,
			Name:      jm.Name,
			StartTime: jm.StartTime,
			Location:  jm.Location,
			Notes:     jm.Notes,
		}
		response = append(response, resp)
	}
	return response, nil
}

func (u user) ActiveJam(useID string) (models.Jam, error) {

	var jam models.Jam
	db := db.NewDB()
	defer db.Close()
	jc := db.JamCollection()
	err := jc.Find(bson.M{"user_id": useID, "current": true}).One(&jam)
	if err != nil {
		return jam, err
	}

	return jam, nil
}
