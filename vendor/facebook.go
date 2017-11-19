package vendor

import (
	"dsound/models"
	"dsound/types"

	fb "github.com/huandu/facebook"
	"gopkg.in/mgo.v2/bson"
)

// FBFetchUser func, fetcher an user from facebook
func FBFetchUser(para types.CreateUserParams) (models.User, error) {

	fbUserID := para.FBID
	accessToken := para.AccessToken

	resp, err := fb.Get("v2.9/me", fb.Params{
		"scopes":       "public_profile",
		"fields":       "first_name,last_name,email",
		"access_token": accessToken,
	})

	if err == nil {
		var firstName string
		var lastName string
		var fbEmail string
		resp.DecodeField("first_name", &firstName)
		resp.DecodeField("last_name", &lastName)
		resp.DecodeField("email", &fbEmail)

		user := models.User{
			ID:        bson.NewObjectId().Hex(),
			FirstName: firstName,
			LastName:  lastName,
			FBID:      fbUserID,
			FBEmail:   fbEmail,
		}
		return user, nil
	}

	return models.User{}, err
}
