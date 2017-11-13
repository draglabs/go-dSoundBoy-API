package vendor

import (
	"dsound/models"
	"dsound/types"

	fb "github.com/huandu/facebook"
	"gopkg.in/mgo.v2/bson"
)

// FBFetchUser func, fetcher an user from facebook
func FBFetchUser(para types.UserRequestParams) (models.User, error) {

	fbUserID := para.FBID
	accessToken := para.AccessToken

	resp, err := fb.Get("/"+fbUserID, fb.Params{
		"fields":       "first_name,last_name",
		"access_token": accessToken,
	})

	if err == nil {
		var firstName string
		var lastName string
		resp.DecodeField("first_name", &firstName)
		resp.DecodeField("last_name", &lastName)
		user := models.User{
			ID:        bson.NewObjectId(),
			FirstName: firstName,
			LastName:  lastName,
			FBID:      fbUserID,
		}
		return user, nil
	}

	return models.User{}, err
}
