package exclusion

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// URL For exclusion list
	URL = "http://private-1de182-mamtrialrankingadjustments4.apiary-mock.com/exclusions"
)

// Entity  struct
type Entity struct {
	Host          string `json:"host"`
	ExcludedSince string `json:"excludedSince"`
	ExcludedTill  string `json:"excludedTill"`
}

// List of Entity
type List []Entity

// GetList - get latest data from exclusion source
func GetList() (List, error) {
	var result List
	resp, err := http.Get(URL)
	if err != nil {
		return result, err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return result, err
	}
	return result, nil
}

// GetExclusionBsonList - get list without exclusion
func GetExclusionBsonList(listExclusion List) ([]bson.M, error) {
	bsonList := []bson.M{}
	for _, exclusionItem := range listExclusion {
		website := exclusionItem.Host
		if exclusionItem.ExcludedSince != "" {
			startDate, _ := time.Parse("2006-01-02", exclusionItem.ExcludedSince)
			endDate := time.Now()
			if exclusionItem.ExcludedTill != "" {
				endDate, _ = time.Parse("2006-01-02", exclusionItem.ExcludedTill)
			}
			bsonList = append(bsonList, bson.M{
				"website": bson.M{
					"$not": bson.RegEx{`/` + website + `$/i`, ""},
				},
				"date": bson.M{
					"$not": bson.M{
						"$gte": startDate,
						"$lt":  endDate,
					},
				},
			})
		}
	}
	return bsonList, nil
}
