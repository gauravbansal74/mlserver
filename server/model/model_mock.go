package model

import "github.com/gauravbansal74/mlserver/pkg/database"

// MockLoadSession with MongoDB
func MockLoadSession() (MongoSession, string) {
	mongoDbInfo := database.LoadConfig("", "")
	session := mongoDbInfo.Init()
	return MongoSession{session.Session}, database.ReadMongoConfig().Database
}
