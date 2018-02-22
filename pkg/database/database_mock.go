package database

// MockLoadConfig - load mongo config
func MockLoadConfig(url, db string) MongoDBInfo {
	return MongoDBInfo{}
}

// MockDail mgo Dail
func (m MongoDBInfo) MockDail(url string) (MongoSession, error) {
	mockSession := MongoSession{}
	return mockSession, nil
}
