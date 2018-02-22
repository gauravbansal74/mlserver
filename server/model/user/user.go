package user

import (
	"encoding/json"
	"errors"

	"gopkg.in/mgo.v2/bson"

	"github.com/gauravbansal74/mlserver/pkg/utils"
	"github.com/gauravbansal74/mlserver/server/model"
)

var (
	errNoResult      = errors.New("no result")
	errAlreadyExists = errors.New("already registered")
)

const (
	collectionName = "users"
)

// Entity website visits entity
type Entity struct {
	ID       bson.ObjectId `bson:"_id" json:"_id"`
	Email    string        `bson:"email" json:"email"`
	Password string        `bson:"-" json:"password"`
	Hash     string        `bson:"hash" json:"-"`
	Salt     string        `bson:"salt" json:"-"`
}

// List of Entity
type List []Entity

// Body read body data and get into struct value
func Body(body []byte) (*Entity, error) {
	entity := &Entity{}
	err := json.Unmarshal(body, &entity)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

// New read body data and get into struct value
func New(body []byte) (*Entity, error) {
	entity := &Entity{}
	err := json.Unmarshal(body, &entity)
	if err != nil {
		return entity, err
	}
	// Set the default parameters
	entity.ID = bson.NewObjectId()
	entity.Salt = bson.NewObjectId().Hex()
	entity.Hash, err = utils.GetHash(entity.Password, entity.Salt)
	if err != nil {
		return entity, err
	}
	return entity, nil
}

// Create create user data
func (v *Entity) Create() error {
	_, err := readOneByField("email", v.Email)
	if err != nil && err.Error() == "not found" {
		session, databaseName := model.LoadSession()
		defer session.Close()
		c := session.DB(databaseName).C(collectionName)
		err := c.Insert(v)
		return err
	}
	if err != nil {
		return err
	}
	return errAlreadyExists
}

// Delete record from db
func (v *Entity) Delete() error {
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	err := c.Remove(bson.M{"email": v.Email})
	return err
}

// GetByEmail - get user data by email
func GetByEmail(email string) (*Entity, error) {
	result, err := readOneByField("email", email)
	return result, err
}

// readOneByField returns the entity that matches the field value
// If no result, it will return ErrNoResult
func readOneByField(name string, value string) (*Entity, error) {
	result := &Entity{}
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	err := c.Find(bson.M{name: value}).One(&result)
	return result, err
}
