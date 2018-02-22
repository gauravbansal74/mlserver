package database

import (
	"testing"
)

func TestLoadConfig(t *testing.T) {
	type args struct {
		url string
		db  string
	}
	tests := []struct {
		name string
		args args
		want *MongoDBInfo
	}{
		{
			name: "config test",
			args: args{
				url: "",
				db:  "",
			},
			want: &MongoDBInfo{
				URL:      "",
				Database: "",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LoadConfig(tt.args.url, tt.args.db)
			if got.Database != tt.want.Database {
				t.Errorf("LoadConfig() = %v, want %v", got.Database, tt.want.Database)
			}
			if got.URL != tt.want.URL {
				t.Errorf("LoadConfig() = %v, want %v", got.URL, tt.want.URL)
			}
		})
	}
}

func TestMongoDBInfo_Init(t *testing.T) {
	mongoDbInfo := MockLoadConfig("", "")
	msession := mongoDbInfo.Init()
	err := msession.Ping()
	if err != nil {
		t.Errorf("Init() = %v", err.Error())
	}
}

func TestMongoDBInfo_IsMongoConnected(t *testing.T) {
	mongoDbInfo := MockLoadConfig("", "")
	msession := mongoDbInfo.IsMongoConnected()
	err := msession.Ping()
	if err != nil {
		t.Errorf("Init() = %v", err.Error())
	}
}
