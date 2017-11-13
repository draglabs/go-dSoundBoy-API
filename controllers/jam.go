package controllers

import (
	"dsound/db"
	"dsound/models"
	"dsound/types"
	"dsound/utils"
	"dsound/vendor"
	"errors"

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

	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	jam := models.Jam{
		UserID:      p.UserID,
		Pin:         utils.GeneratePin(4),
		ID:          bson.NewObjectId(),
		Name:        p.Name,
		Location:    p.Location,
		Coordinates: []float64{p.Lat, p.Lng},
	}
	err := c.Insert(jam)
	if err == nil {
		go User.UpdateCurrentJam(p.UserID, jam.ID.String())
		return jam, nil
	}

	return jam, err
}

func (j jam) Upload(p types.UploadJamParams) error {
	s3URL, err := vendor.UploadToS3(p.TempFileURL, p.UserID)
	if err != nil {
		go vendor.CleanupAfterUpload(p.TempFileURL)
		return err
	}
	go vendor.CleanupAfterUpload(p.TempFileURL)
	recording := models.Recordings{
		FileName:  p.FileName,
		JamID:     p.JamID,
		StartTime: p.StartTime,
		EndTime:   p.EndTime,
		S3url:     s3URL,
	}
	go updateRecordings(p.JamID, recording)
	return nil
}
func (j jam) Join(p types.JoinJamRequestParams) (types.JamResponse, error) {

	if jm, err := findByPin(p.Pin); err == nil {
		go User.UpdateCurrentJam(p.UserID, jm.ID.String())
		go updateCollabators(jm.ID.String(), p.UserID)
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

func (j jam) FindById(id string) (models.Jam, error) {
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

func updateCollabators(id, userID string) {
	var jm models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	usr, _ := User.FindByID(userID)
	if err := c.FindId(bson.ObjectIdHex(id)).One(&jm); err == nil {
		collabs := jm.Collaborators
		collabs = append(collabs, usr)
		er := c.Update(jm, collabs)
		println(er)
	}

}
func updateRecordings(jamID string, r models.Recordings) {
	var jm models.Jam
	db := db.NewDB()
	defer db.Close()
	c := db.JamCollection()
	if err := c.FindId(bson.ObjectIdHex(jamID)).One(&jm); err == nil {
		recordings := jm.Recordings
		recordings = append(recordings, r)
		er := c.Update(jm, recordings)
		println(er)
	}

}
