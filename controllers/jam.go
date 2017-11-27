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

// Create func, creates a new jam
// and  updates the current jam
// on the user
func (j jam) Create(p types.JamRequestParams) (models.Jam, error) {
	j.UpdateActiveJam(p.UserID)
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	jam := models.Jam{
		IsCurrent:   true,
		ID:          bson.NewObjectId().Hex(),
		UserID:      p.UserID,
		StartTime:   time.Now().String(),
		Pin:         utils.GeneratePin(4),
		Name:        p.Name,
		Location:    p.Location,
		Coordinates: []float64{p.Lat, p.Lng},
	}
	err := c.Insert(jam)
	if err == nil {
		go User.UpdateCurrentJam(p.UserID, jam)

		return jam, nil
	}

	return jam, err
}

func (j jam) Upload(p types.UploadJamParams) error {
	id := bson.NewObjectId().Hex()
	s3URL, err := vendor.UploadToS3(p.TempFileURL, p.UserID, id)
	if err != nil {
		go vendor.CleanupAfterUpload(p.TempFileURL)
		return err
	}
	go vendor.CleanupAfterUpload(p.TempFileURL)
	recording := models.Recordings{
		ID:        id,
		UserID:    p.UserID,
		FileName:  p.FileName, // not in use, not sent from client
		JamID:     p.JamID,
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		S3url:     s3URL,
	}
	go createRecording(recording)
	return nil
}
func (j jam) Join(p types.JoinJamRequestParams) (types.JamResponse, error) {

	if jm, err := findByPin(p.Pin); err == nil {
		j.UpdateActiveJam(p.UserID)
		go User.UpdateCurrentJam(p.UserID, jm)
		go updateCollabators(jm.ID, p.UserID)
		return types.JamResponse{
			ID:        jm.ID,
			Name:      jm.Name,
			StartTime: jm.StartTime,
			Location:  jm.Location,
			Notes:     jm.Notes,
		}, nil
	}
	return types.JamResponse{}, errors.New("unable to join")
}

// Update, updates the jam fields
func (j jam) Update(p types.UpdateJamRequestParams) (models.Jam, error) {
	var jam models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	err := c.Update(bson.M{"_id": p.ID}, bson.M{"$set": bson.M{"name": p.Name, "location": p.Location, "notes": p.Notes}})
	if err != nil {
		fmt.Println("updating", err)
		return jam, err
	}
	err = c.Find(bson.M{"_id": p.ID}).One(&jam)
	fmt.Println("fetching after update", err)
	return jam, err
}

// UpdateActiveJam updates the current jam from
// being active to inactive.
func (j jam) UpdateActiveJam(userID string) {
	var activeJam models.Jam
	db := db.NewDB()
	defer db.Close()
	err := db.JamCollection().Find(bson.M{"user_id": userID, "is_current": true}).One(&activeJam)
	if err == nil {
		err = db.JamCollection().Update(bson.M{"_id": activeJam.ID}, bson.M{"$set": bson.M{"is_current": false, "end_time": time.Now().String()}})
	}

}

// FindById finds a jam by id
func (j jam) FindByID(id string) (models.Jam, error) {
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

func findByPin(pin string) (models.Jam, error) {
	var jm models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	err := c.Find(bson.M{"pin": pin}).One(&jm)
	if err == nil {
		return jm, nil
	}
	return jm, err
}

// Recordings func, will fetch all the recordings for a jam
func Recordings(jamID string) ([]models.Recordings, error) {
	var recordings []models.Recordings
	db := db.NewDB()
	defer db.Close()
	err := db.RecordingsCollection().Find(bson.M{"jam_id": jamID}).All(&recordings)
	return recordings, err
}

func updateCollabators(jamID, userID string) {
	var jm models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	usr, _ := User.FindByID(userID)
	if err := c.FindId(jamID).One(&jm); err == nil {
		collabs := jm.Collaborators
		collabs = append(collabs, usr)
		er := c.Update(bson.M{"_id": jamID}, bson.M{"$set": bson.M{"collaborators": collabs}})
		fmt.Println(er)
	}

}
func createRecording(r models.Recordings) error {
	db := db.NewDB()
	defer db.Close()
	err := db.RecordingsCollection().Insert(r)
	if err != nil {
		fmt.Println("error creating recording ", err)
	}

	return err
}
func (j jam) Details(id string) (models.Jam, error) {
	recordings, err := Recordings(id)
	jam, err := j.FindByID(id)
	jam.Recordings = recordings
	return jam, err
}
