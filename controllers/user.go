package controllers

import (
	"dsound/db"
	"dsound/models"
)

type User struct {
}

func (u User) find(id string) (models.User, error) {
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
