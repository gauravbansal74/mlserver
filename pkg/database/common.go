package database

// DatabaseCommon interface
type DatabaseCommon interface {
	Dial(url string) (MongoSession, error)
}
