package exclusion

import (
	"testing"
)

func TestGetList(t *testing.T) {
	_, err := GetList()
	if err != nil {
		t.Errorf("GetList() got error %v", err.Error())
	}
}

func TestGetExclusionBsonList(t *testing.T) {
	list, err := GetList()
	if err != nil {
		t.Errorf("GetList() got error %v", err.Error())
	}
	if len(list) > 0 {
		bsonList, err := GetExclusionBsonList(list)
		if err != nil {
			t.Errorf("GetExclusionBsonList() got error %v", err.Error())
		}
		if len(bsonList) != len(list) {
			t.Errorf("GetExclusionBsonList() error while creating query")
		}
	}
}
