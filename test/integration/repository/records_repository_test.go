package service_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/ahmetcanaydemir/go-rest/pkg/api/repository"
	"github.com/ahmetcanaydemir/go-rest/pkg/configs"
	"github.com/ahmetcanaydemir/go-rest/pkg/model"
)

func Test_RecordsRepository(t *testing.T) {
	var tests = []struct {
		name             string
		mongoRequestJson string
		resultExpected   bool
	}{
		{"exists request", `{"startDate":"2016-01-26","endDate":"2018-02-02","minCount":10,"maxCount":10000}`, true},
		{"no record request", `{"startDate":"3301-26","endDate":"2018-02-02","minCount":99998,"maxCount":99999}`, false},
	}

	connStr := os.Getenv("MONGO_URI")

	if connStr == "" {
		t.Errorf("required environment variable MONGO_URI not set")
		return
	}
	configs.Server.Config.DbConnectionString = connStr

	recordsRepository := repository.NewRecordsRepository()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var request model.MongoRequest
			err := json.Unmarshal([]byte(tt.mongoRequestJson), &request)
			if err != nil {
				t.Errorf("GetMongoController(%s) unexpected error occured during json unmarshal: %s", tt.name, err.Error())
			}

			gotResponse, err := recordsRepository.GetRecords(request)
			if err != nil {
				t.Errorf("GetMongoController(%s) unexpected error occured: %s", tt.name, err.Error())
			}
			if (tt.resultExpected && len(gotResponse) == 0) || (!tt.resultExpected && len(gotResponse) != 0) {
				t.Errorf("GetMongoController(%s) is result expected: %v, result count: %v", tt.name, tt.resultExpected, len(gotResponse))
			}
		})
	}
}
