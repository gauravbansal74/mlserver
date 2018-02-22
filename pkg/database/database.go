package database

import (
	"log"

	"github.com/gauravbansal74/mlserver/pkg/logger"
	mgo "gopkg.in/mgo.v2"
)

var (
	// Mongo wrapper
	Mongo          MongoSession
	mongoDatabases MongoDBInfo
)

// MongoSession is currently a Mongo session.
type MongoSession struct {
	*mgo.Session
}

// Database interface
type Database interface {
	LoadConfig(url, db string) MongoDBInfo
	Init() MongoSession
	IsMongoConnected() MongoSession
	ReadMongoConfig() MongoSession
}

// MongoDBInfo is the details for the database connection
type MongoDBInfo struct {
	URL      string
	Database string
}

// LoadConfig - load mongo config
func LoadConfig(url, db string) MongoDBInfo {
	mongoDatabases = MongoDBInfo{
		URL:      url,
		Database: db,
	}
	return mongoDatabases
}

// Init Connect to the database
func (m MongoDBInfo) Init() MongoSession {
	mongoDatabases = MongoDBInfo{
		URL:      m.URL,
		Database: m.Database,
	}
	mSession, err := mgo.Dial(mongoDatabases.URL)
	if err != nil {
		log.Fatal(err.Error(), "Error while connecting to MongoDB")
	}
	Mongo = MongoSession{mSession}
	Mongo.SetMode(mgo.Monotonic, true)
	logger.Info("Successfully Connected to MongoDB")
	return Mongo
}

// IsMongoConnected returns true if MongoDB is available
func (m MongoDBInfo) IsMongoConnected() MongoSession {
	err := Mongo.Ping()
	if err != nil {
		Mongo = m.Init()
	}
	return Mongo
}

// ReadMongoConfig returns the database information
func ReadMongoConfig() MongoDBInfo {
	return mongoDatabases
}
