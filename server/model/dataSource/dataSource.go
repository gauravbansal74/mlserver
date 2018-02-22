package dataSource

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gauravbansal74/mlserver/server/model"
)

const (
	collectionName = "data_files"
)

// Entity website visits entity
type Entity struct {
	ID        bson.ObjectId `bson:"_id" json:"_id"`
	Name      string        `bson:"fileName" json:"fileName"`
	Message   string        `bson:"message" json:"message"`
	Status    string        `bson:"status" json:"status"`
	CreatedAt time.Time     `bson:"created" json:"created"`
}

// List of Entity
type List []Entity

// Create create entry for website visits
func (v *Entity) Create() error {
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	err := c.Insert(v)
	return err
}

// GetAllRecords - get website visits data by date
func GetAllRecords() (List, error) {
	var result List
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	err := c.Find(bson.M{
		"created": bson.M{
			"$lt": time.Now(),
		},
	}).Sort("-created").All(&result)
	return result, err
}
