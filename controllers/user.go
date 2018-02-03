package controllers

import (
	"dsound/db"
	"dsound/models"
	"dsound/types"
	"dsound/vendor"
	"fmt"

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
	err := c.Find(bson.M{"_id": id}).One(&user)
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
func (u user) UpdateCurrentJam(userID string, jam models.Jam) error {

	db := db.NewDB()
	defer db.Close()
	c := db.UserCollection()

	err := c.Update(bson.M{"_id": userID}, bson.M{"$set": bson.M{"current_jam": jam}})
	if err != nil {
		fmt.Println("user active jam not found", err)
		return err
	}

	return err
}

func (u user) Activity(userID string) ([]types.JamResponse, error) {
	var response []types.JamResponse
	var jams []models.Jam
	db := db.NewDB()
	defer db.Close()
	jc := db.JamCollection()
	err := jc.Find(bson.M{"user_id": userID}).All(&jams)
	if err != nil {
		return response, err
	}
	for _, jm := range jams {
		count := len(jm.Collaborators)
		resp := types.JamResponse{
			ID:            jm.ID,
			Name:          jm.Name,
			StartTime:     jm.StartTime,
			EndtTime:      jm.EndTime,
			Location:      jm.Location,
			Notes:         jm.Notes,
			Collaborators: count,
			Link:          jm.Link,
			Coordinates:   jm.Coordinates,
		}
		response = append(response, resp)
	}
	return response, nil
}

func extractJamFromCollaboration() {

}
func (u user) ActiveJam(useID string) (models.Jam, error) {
	var jam models.Jam
	db := db.NewDB()
	defer db.Close()
	jc := db.JamCollection()
	err := jc.Find(bson.M{"user_id": useID, "is_current": true}).One(&jam)
	return jam, err

}
func (u user) Update(userID string) (models.User, error) {
	var user models.User
	store := db.NewDB()
	defer store.Close()
	err := store.UserCollection().FindId(userID).One(&user)
	if err != nil {
		return user, err
	}
	return user, nil
}
