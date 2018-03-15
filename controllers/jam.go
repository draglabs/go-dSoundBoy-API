package controllers

import (
	"dsound/db"
	"dsound/models"
	"dsound/types"
	"dsound/utils"
	"dsound/vendor"
	"errors"
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type jam struct {
}

func newJam() jam {
	return jam{}
}

var Jam = newJam()

// Creates a new jam and updates the current jam on the user
//
// Param: p, an object of JamRequestParams
//
// Returns: a Jam response type and an error if something went wrong
func (j jam) Create(p types.JamRequestParams) (models.Jam, error) {
	j.UpdateActiveJam(p.UserID)                                                     // updates the active jam on the user from the parameter
	db := db.NewDB()                                                                // creates an instance of the database
	defer db.Close()                                                                // closes the database once the function returns
	c := db.JamCollection()                                                         // creates an object referencing a jam location in the database
	jam := models.Jam {                                                             // instantiates a jam object
		IsCurrent:   true,                                                              // is current jam
		ID:          bson.NewObjectId().Hex(),                                          // jam id is new hexadecimal object id
		UserID:      p.UserID,                                                          // user id comes from parameter
		StartTime:   time.Now().String(),                                               // start time is current system time as string
		Pin:         utils.GeneratePin(4),                                         // generates the pin with max length of 4
		Name:        p.Name,                                                            // jam name comes from parameter
		Location:    p.Location,                                                        // jam location comes from parameter
		Coordinates: []float64{p.Lat, p.Lng},                                           // coordinates come from parameters and are set as 64-bit floats
	}
	err := c.Insert(jam)                                                            // returns and instantiates error from inserting the jam into the database's jam collection
	if err == nil {                                                                 // if there is no error
		go User.UpdateCurrentJam(p.UserID, jam)                                     // update the users's current jam with this one

		return jam, nil                                                             // return the current created jam and no error
	}

	return jam, err                                                                 // if there was an error, return it along with the jam
}

// Uploads current jam to S3
//
// Param: p, an object of UploadJamParams type
//
// Returns: error, if something went wrong
func (j jam) Upload(p types.UploadJamParams) error {
	id := bson.NewObjectId().Hex()                                                  // creates new hexadecimal object id for the recording
	s3URL, err := vendor.UploadToS3(p.TempFileURL, p.UserID, id)                    // passes into the uploader a temporary local file url, the user id, and the object id, returns a new S3 url
	if err != nil {                                                                 // if there is an error
		go vendor.CleanupAfterUpload("temp")                                  // delete the temporary local storage folder as a goroutine
		return err                                                                  // return the error
	}
	go vendor.CleanupAfterUpload("temp")                                      // delete the temporary local storage folder as a goroutine
	recording := models.Recordings{                                                 // instantiate a recording object representing a blank Recording models object
		ID:        id,
		UserID:    p.UserID,
		FileName:  p.FileName, // not in use, not sent from client
		JamID:     p.JamID,
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		S3url:     s3URL,
	}
	go createRecording(recording)                                                   // assign the recording object data from the createRecording function as a goroutine
	return err                                                                      // return an error even though it succeeded
}

// Joins a user to a jam by referencing jam's unique pin code. If the pin is invalid, the api is unable to join the user to the jam.
//
// Param: p, an object of JoinJamRequestParams
//
// Returns: a JamResponse type and an error if something went wrong
func (j jam) Join(p types.JoinJamRequestParams) (types.JamResponse, error) {
	jm, err := findByPin(p.Pin)                                                     // finds the jam by pin code from the parameter
	if err != nil {                                                                 // if there is an error
		fmt.Println("jam not found with pin ", p.Pin)
		return types.JamResponse{}, err                                             // returns an error and a blank Jam Response object
	} else {                                                                        // if there is no error
		fmt.Println("jam found name:", jm.Name)
		User.UpdateCurrentJam(p.UserID, jm)                                         // updates the current jam by taking in parameter's user id and the current jam by pin code
		j.UpdateActiveJam(p.UserID)                                                 // updates the user's active jam to this one

		updateCollaborators(jm.ID, p.UserID)                                        // adds the user to the list of collaborators
		return types.JamResponse{                                                   // returns a new Jam Response object with information from the jam found by pin code
			ID:        jm.ID,
			Name:      jm.Name,
			StartTime: jm.StartTime,
			Location:  jm.Location,
			Notes:     jm.Notes,
			Pin:       jm.Pin,
		}, nil
	}
	return types.JamResponse{}, errors.New("unable to join")                   // if the pin doesn't exist in a jam, api is unable to join the jam
}

// Updates the jam fields. Get information to update, updates the copy on the database, and then returns the newly re-fetched object.
//
// Param: p, an object of UpdateJamRequestParams
//
// Returns: a JamResponse type and an error if something went wrong
func (j jam) Update(p types.UpdateJamRequestParams) (types.JamResponse, error) {
	var jam models.Jam                                                              // instantiates an object representing a jam
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	c := db.JamCollection()                                                         // gets a jam collection from the database
	err := c.Update(bson.M{"_id": p.ID}, bson.M{"$set": bson.M{"name": p.Name, "location": p.Location, "notes": p.Notes}}) // updates the jam collection with information from the parameter
	if err != nil {                                                                 // if there is an error
		fmt.Println("updating", err)
		return types.JamResponse{}, err                                             // returns an empty Jam Response object with an error
	}
	err = c.Find(bson.M{"_id": p.ID}).One(&jam)                                     // sets the current jam object to the one found by the parameter's id
	fmt.Println("fetching after update", err)
	return types.JamResponse{                                                       // returns a new Jam Response object with information from the jam as updated
		ID:        jam.ID,
		Name:      jam.Name,
		StartTime: jam.StartTime,
		Location:  jam.Location,
		Notes:     jam.Notes,
		Link:      jam.Link,
	}, err
}

// UpdateActiveJam updates the current jam from being active to inactive.
//
// Param: a user ID
//
// Returns: an error if something went wrong
func (j jam) UpdateActiveJam(userID string) error {
	var activeJam models.Jam                                                        // instantiates an object representing a jam
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	err := db.JamCollection().Find(bson.M{"user_id": userID, "is_current": true}).One(&activeJam) // sets the active jam to the one found in the database's jam collection
	if err != nil {                                                                 // if there is an error
		fmt.Println("error finding jam")
		return err
	} else {                                                                        // if there is no error
		err = db.JamCollection().Update(bson.M{"_id": activeJam.ID}, bson.M{"$set": bson.M{"is_current": false, "end_time": time.Now().String()}}) // reset the error to the updated jam
		return err
	}
	if err != nil {                                                                 // if there is a new error from updating the jam
		fmt.Println("error updating active jam ", err)
		return err
	}
	return nil                                                                      // necessary returns already occurred
}

// Finds a jam by id
//
// Param: a jam ID
//
// Returns: a jam model and an error if something went wrong
func (j jam) FindByID(id string) (models.Jam, error) {
	var jm models.Jam                                                               // instantiates an object representing a jam
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	c := db.JamCollection()                                                         // instantiates a variable representing a jam collection inside the database
	err := c.FindId(id).One(&jm)                                                    // looks through the jam collection to find the jam by its id
	if err == nil {                                                                 // if there is no error
		return jm, nil                                                              // returns the jam and no error
	}
	return jm, err                                                                  // returns the empty jam and an error since something went wrong
}

// Finds a jam by pin
//
// Param: a jam pin
//
// Returns: a jam model and an error if something went wrong
func findByPin(pin string) (models.Jam, error) {
	var jm models.Jam                                                               // instantiates an object representing a jam
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	c := db.JamCollection()                                                         // instantiates a variable representing a jam collection in the database
	err := c.Find(bson.M{"pin": pin}).One(&jm)                                      // assigns the jam object one that is found in the jam collection by pin
	if err != nil {                                                                 // if there is an error
		return jm, err
	}
	return jm, nil
}

// Recordings func, will fetch all the recordings for a jam
//
// Param: a jam ID
//
// Returns: an array of Recording models and an error if something went wrong
func Recordings(jamID string) ([]models.Recordings, error) {
	var recordings []models.Recordings                                              // instantiates an array of Recording models
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	err := db.RecordingsCollection().Find(bson.M{"jam_id": jamID}).All(&recordings) // sets the recordings to the ones found by jam ID and the error from the Find operation
	return recordings, err
}

// Adds a user to a jam's collaborators
//
// Param: a jam ID and a user ID
//
// Returns: nothing
func updateCollaborators(jamID, userID string) {
	var jm models.Jam                                                               // instantiates a jam model
	db := db.NewDB()                                                                // instantiates a new database connection
	defer db.Close()                                                                // closes the database connection once the function returns
	c := db.JamCollection()                                                         // instantiates c to the jam collection in the database
	usr, _ := User.FindByID(userID)                                                 // fetches the user by the user's id
	if err := c.Find(bson.M{"_id": jamID}).One(&jm); err == nil {                   // if there was no error while fetching the jam by its id
		collabs := jm.Collaborators                                                 // sets the jam model's collaborators to the ones found by fetching the jam
		collabs = append(collabs, usr)                                              // adds the user to the collaborators
		er := c.Update(bson.M{"_id": jamID}, bson.M{"$set": bson.M{"collaborators": collabs}}) // replaces the copy on the database with this new updated one
		fmt.Println("error updating collabators ", er)
	}

}

// Creates a recording from a recording model
//
// Param: r, a Recording model
//
// Returns: an error if something went wrong
func createRecording(r models.Recordings) error {
	db := db.NewDB()                                                                // instantiates new database connection
	defer db.Close()                                                                // closes it when the function returns
	err := db.RecordingsCollection().Insert(r)                                      // inserts the parameter into the collection of recordings
	if err != nil {                                                                 // if there was an error
		fmt.Println("error creating recording ", err)
	}

	return err
}

// Gets the recordings' details with a jam ID
//
// Param: id, a jam ID
//
// Returns: a Jam model and an error if something went wrong
func (j jam) Details(id string) (models.Jam, error) {
	recordings, err := Recordings(id)                                               // fetches the recordings by their id
	jam, err := j.FindByID(id)                                                      // finds the jam by jam id
	jam.Recordings = recordings                                                     // sets the jam's recordings with the ones ones instantiated here
	return jam, err
}
