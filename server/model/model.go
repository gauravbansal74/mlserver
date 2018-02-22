package model

import (
	mgo "gopkg.in/mgo.v2"

	"github.com/gauravbansal74/mlserver/pkg/database"
)

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// LoadSession with MongoDB
func LoadSession() (MongoSession, string) {
	dbInfo := database.ReadMongoConfig()
	mongo := dbInfo.IsMongoConnected()
	session := mongo.Copy()
	return MongoSession{session}, database.ReadMongoConfig().Database
}

// MongoCollection wraps a mgo.Collection to embed methods in models.
type MongoCollection struct {
	*mgo.Collection
}

// MongoQuery wraps a mgo.Query to embed methods in models.
type MongoQuery struct {
	*mgo.Query
}

// MongoDatabase wraps a mgo.Database to embed methods in models.
type MongoDatabase struct {
	*mgo.Database
}

// DataLayer is an interface to access to the database struct.
type DataLayer interface {
	C(name string) Collection
}

// Session is an interface to access to the Session struct.
type Session interface {
	DB(name string) DataLayer
	Close()
}

// Collection is an interface to access to the collection struct.
type Collection interface {
	Find(query interface{}) Query
	Pipe(pipeline interface{}) Pipe
	Count() (n int, err error)
	Insert(docs ...interface{}) error
	Remove(selector interface{}) error
	Update(selector interface{}, update interface{}) error
}

// Query is an interface to access to the database struct
type Query interface {
	All(result interface{}) error
	One(result interface{}) (err error)
	Sort(fields ...string) MongoQuery
	Distinct(key string, result interface{}) error
}

// Pipe is an interface to access to the database struct
type Pipe interface {
	All(result interface{}) error
}
