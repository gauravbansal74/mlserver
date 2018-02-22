package visits

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/gauravbansal74/mlserver/server/exclusion"
	"github.com/gauravbansal74/mlserver/server/model"
)

const (
	collectionName = "website_visits"
)

// Entity website visits entity
type Entity struct {
	ID      bson.ObjectId `bson:"_id" json:"_id"`
	Date    time.Time     `bson:"date" json:"date"`
	Website string        `bson:"website" json:"website"`
	Total   int64         `bson:"visits" json:"visits"`
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

// GetOverAll - Get overall total website visits
func GetOverAll(listExclusion exclusion.List) ([]interface{}, error) {
	var result []interface{}
	exclusionList, err := exclusion.GetExclusionBsonList(listExclusion)
	if err != nil {
		return result, err
	}
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"$and": exclusionList,
			},
		},
		bson.M{"$group": bson.M{
			"_id": "$website",
			"visits": bson.M{
				"$sum": "$visits",
			},
		}},
		bson.M{
			"$sort": bson.M{
				"visits": -1,
			},
		},
		bson.M{
			"$limit": 5,
		},
	}
	pipe := c.Pipe(pipeline)
	err = pipe.All(&result)
	return result, err
}

// GetByDate - get website visits data by date
func GetByDate(selectedDate time.Time, listExclusion exclusion.List) ([]interface{}, error) {
	var result []interface{}
	exclusionList, err := exclusion.GetExclusionBsonList(listExclusion)
	if err != nil {
		return result, err
	}
	session, databaseName := model.LoadSession()
	defer session.Close()
	c := session.DB(databaseName).C(collectionName)
	pipeline := []bson.M{
		bson.M{
			"$match": bson.M{
				"$and": []bson.M{
					bson.M{
						"date": bson.M{
							"$gte": selectedDate,
							"$lt":  selectedDate.Add(time.Minute * 60 * 24),
						},
					},
					bson.M{
						"$and": exclusionList,
					},
				},
			},
		},
		bson.M{"$group": bson.M{
			"_id": "$website",
			"visits": bson.M{
				"$sum": "$visits",
			},
		}},
		bson.M{
			"$sort": bson.M{
				"visits": -1,
			},
		},
	}
	pipe := c.Pipe(pipeline)
	err = pipe.All(&result)
	return result, err
}
